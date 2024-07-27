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

	chWaitReply []chan *entry.Entry
	chWaitWg    []sync.WaitGroup
	offReply    bool
	ch          chan *entry.Entry
	chWg        sync.WaitGroup

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
	for i := 0; i < opts.Clients; i++ {
		rw.clients = append(rw.clients, client.NewRedisClient(ctx, opts.Address, opts.Username, opts.Password, opts.Tls, false))
	}
	rw.ch = make(chan *entry.Entry, 1024)
	if opts.OffReply {
		log.Infof("turn off the reply of write")
		rw.offReply = true
		for _, c := range rw.clients {
			c.Send("CLIENT", "REPLY", "OFF")
		}
	} else {
		rw.chWaitReply = make([]chan *entry.Entry, opts.Clients)
		rw.chWaitWg = make([]sync.WaitGroup, opts.Clients)
		for idx := range rw.chWaitReply {
			rw.chWaitReply[idx] = make(chan *entry.Entry, config.Opt.Advanced.PipelineCountLimit)
			rw.chWaitWg[idx].Add(1)
		}
		for idx := range rw.clients {
			go rw.processReply(idx)
		}
	}
	return rw
}

func (w *redisStandaloneWriter) Close() {
	if !w.offReply {
		close(w.ch)
		w.chWg.Wait()
		for _, ch := range w.chWaitReply {
			close(ch)
		}
		for idx := range w.chWaitWg {
			w.chWaitWg[idx].Wait()
		}
	}
}

func (w *redisStandaloneWriter) StartWrite(ctx context.Context) chan *entry.Entry {
	w.chWg = sync.WaitGroup{}
	w.chWg.Add(1)
	go func() {
		for e := range w.ch {
			// switch db if we need
			if w.DbId != e.DbId {
				w.switchDbTo(e.DbId)
			}
			// send
			bytes := e.Serialize()
			for e.SerializedSize+atomic.LoadInt64(&w.stat.UnansweredBytes) > config.Opt.Advanced.TargetRedisClientMaxQuerybufLen {
				time.Sleep(1 * time.Nanosecond)
			}
			log.Debugf("[%s] send cmd. cmd=[%s]", w.stat.Name, e.String())
			// random select a client to send
			rand.New(rand.NewSource(time.Now().UnixNano()))
			selected := rand.Intn(len(w.clients))
			if !w.offReply {
				w.chWaitReply[selected] <- e
				atomic.AddInt64(&w.stat.UnansweredBytes, e.SerializedSize)
				atomic.AddInt64(&w.stat.UnansweredEntries, 1)
			}
			w.clients[selected].SendBytes(bytes)
		}
		w.chWg.Done()
	}()

	return w.ch
}

func (w *redisStandaloneWriter) Write(e *entry.Entry) {
	w.ch <- e
}

func (w *redisStandaloneWriter) switchDbTo(newDbId int) {
	log.Debugf("[%s] switch db to [%d]", w.stat.Name, newDbId)
	w.DbId = newDbId
	for idx, c := range w.clients {
		c.Send("select", strconv.Itoa(newDbId))
		if !w.offReply {
			w.chWaitReply[idx] <- &entry.Entry{
				Argv:    []string{"select", strconv.Itoa(newDbId)},
				CmdName: "select",
			}
		}
	}
}

func (w *redisStandaloneWriter) processReply(idx int) {
	for e := range w.chWaitReply[idx] {
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
	w.chWaitWg[idx].Done()
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
