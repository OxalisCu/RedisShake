package writer

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"RedisShake/internal/client"
	"RedisShake/internal/client/proto"
	"RedisShake/internal/config"
	"RedisShake/internal/entry"
	"RedisShake/internal/log"
)

type RedisWriterOptions struct {
	Cluster  bool   `mapstructure:"cluster" default:"false"`
	Sentinel bool   `mapstructure:"sentinel" default:"false"`
	Master   string `mapstructure:"master" default:""`
	Address  string `mapstructure:"address" default:""`
	Username string `mapstructure:"username" default:""`
	Password string `mapstructure:"password" default:""`
	Tls      bool   `mapstructure:"tls" default:"false"`
	OffReply bool   `mapstructure:"off_reply" default:"false"`
	Clients  int    `mapstructure:"clients" default:"1"`
}

type redisStandaloneWriter struct {
	address string
	clients []*client.Redis
	DbId    int
	rand    *rand.Rand

	clientNum     int
	offReply      bool
	chWaitReplies []chan *entry.Entry
	chWaitWgs     []sync.WaitGroup
	chs           []chan *entry.Entry
	chWgs         []sync.WaitGroup

	stat struct {
		Name              string `json:"name"`
		UnansweredBytes   int64  `json:"unanswered_bytes"`
		UnansweredEntries int64  `json:"unanswered_entries"`
	}
}

func NewRedisStandaloneWriter(ctx context.Context, opts *RedisWriterOptions) Writer {
	rw := new(redisStandaloneWriter)
	rw.address = opts.Address
	rw.stat.Name = "writer_" + strings.Replace(opts.Address, ":", "_", -1)
	rw.clientNum = opts.Clients
	rw.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < rw.clientNum; i++ {
		rw.clients = append(rw.clients, client.NewRedisClient(ctx, opts.Address, opts.Username, opts.Password, opts.Tls, false))
	}
	rw.chs = make([]chan *entry.Entry, rw.clientNum)
	rw.chWgs = make([]sync.WaitGroup, rw.clientNum)
	for idx := range rw.chs {
		rw.chs[idx] = make(chan *entry.Entry, 1024)
	}
	if opts.OffReply {
		log.Infof("turn off the reply of write")
		rw.offReply = true
		for _, c := range rw.clients {
			c.Send("CLIENT", "REPLY", "OFF")
		}
	} else {
		rw.chWaitReplies = make([]chan *entry.Entry, rw.clientNum)
		rw.chWaitWgs = make([]sync.WaitGroup, rw.clientNum)
		for idx := range rw.chWaitReplies {
			rw.chWaitReplies[idx] = make(chan *entry.Entry, config.Opt.Advanced.PipelineCountLimit)
			rw.chWaitWgs[idx].Add(1)
		}
		go rw.processReply()
	}
	return rw
}

func (w *redisStandaloneWriter) Close() {
	for idx := range w.chs {
		close(w.chs[idx])
	}
	for idx := range w.chWgs {
		w.chWgs[idx].Wait()
	}
	if !w.offReply {
		for idx := range w.chWaitReplies {
			close(w.chWaitReplies[idx])
		}
		for idx := range w.chWaitWgs {
			w.chWaitWgs[idx].Wait()
		}
	}
}

func (w *redisStandaloneWriter) StartWrite(ctx context.Context) (ch chan *entry.Entry) {
	for i := range w.clients {
		w.chWgs[i].Add(1)
		go func(idx int) {
			for e := range w.chs[idx] {
				// switch db if we need
				if w.DbId != e.DbId {
					w.switchDbTo(idx, e.DbId)
				}
				// send
				bytes := e.Serialize()
				for e.SerializedSize+atomic.LoadInt64(&w.stat.UnansweredBytes) > config.Opt.Advanced.TargetRedisClientMaxQuerybufLen {
					time.Sleep(1 * time.Nanosecond)
				}
				log.Debugf("[%s] send cmd. cmd=[%s]", w.stat.Name, e.String())
				if !w.offReply {
					w.chWaitReplies[idx] <- e
					atomic.AddInt64(&w.stat.UnansweredBytes, e.SerializedSize)
					atomic.AddInt64(&w.stat.UnansweredEntries, 1)
				}
				w.clients[idx].SendBytes(bytes)
			}
			w.chWgs[idx].Done()
		}(i)
	}

	return nil
}

func (w *redisStandaloneWriter) Write(e *entry.Entry) {
	// random select a client to send
	selected := w.rand.Intn(len(w.clients))
	w.chs[selected] <- e
}

func (w *redisStandaloneWriter) switchDbTo(idx, newDbId int) {
	log.Debugf("[%s] switch db to [%d]", w.stat.Name, newDbId)
	w.DbId = newDbId
	w.clients[idx].Send("select", strconv.Itoa(newDbId))
	if !w.offReply {
		w.chWaitReplies[idx] <- &entry.Entry{
			Argv:    []string{"select", strconv.Itoa(newDbId)},
			CmdName: "select",
		}
	}
}

func (w *redisStandaloneWriter) processReply() {
	for i := range w.chWaitReplies {
		w.chWaitWgs[i].Add(1)
		go func(idx int) {
			for e := range w.chWaitReplies[idx] {
				reply, err := w.clients[idx].Receive()
				log.Debugf("[%s] receive reply. reply=[%v], cmd=[%s]", w.stat.Name, reply, e.String())

				// It's good to skip the nil error since some write commands will return the null reply. For example,
				// the SET command with NX option will return nil if the key already exists.
				if err != nil && !errors.Is(err, proto.Nil) {
					if err.Error() == "BUSYKEY Target key name already exists." {
						if config.Opt.Advanced.RDBRestoreCommandBehavior == "skip" {
							log.Debugf("[%s] redisStandaloneWriter received BUSYKEY reply. cmd=[%s]", w.stat.Name, e.String())
						} else if config.Opt.Advanced.RDBRestoreCommandBehavior == "panic" {
							log.Panicf("[%s] redisStandaloneWriter received BUSYKEY reply. cmd=[%s]", w.stat.Name, e.String())
						}
					} else {
						log.Panicf("[%s] receive reply failed. cmd=[%s], error=[%v]", w.stat.Name, e.String(), err)
					}
				}
				if strings.EqualFold(e.CmdName, "select") { // skip select command
					continue
				}
				atomic.AddInt64(&w.stat.UnansweredBytes, -e.SerializedSize)
				atomic.AddInt64(&w.stat.UnansweredEntries, -1)
			}
			w.chWaitWgs[idx].Done()
		}(i)
	}
}

func (w *redisStandaloneWriter) Status() interface{} {
	return w.stat
}

func (w *redisStandaloneWriter) StatusString() string {
	return fmt.Sprintf("[%s]: unanswered_entries=%d", w.stat.Name, atomic.LoadInt64(&w.stat.UnansweredEntries))
}

func (w *redisStandaloneWriter) StatusConsistent() bool {
	return atomic.LoadInt64(&w.stat.UnansweredBytes) == 0 && atomic.LoadInt64(&w.stat.UnansweredEntries) == 0
}
