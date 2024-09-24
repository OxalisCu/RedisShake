logdir=/home/duk/code/RedisShake/demo/bench
pprof_port=40000

declare -A mode2filename
mode2filename["profile"]="pprof.redis-shake.samples.cpu.001.pb.gz"
mode2filename["goroutine"]="pprof.goroutine.001.pb.gz"
mode2filename["heap"]="pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz"
# mode2filename["allocs"]="pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz"
# mode2filename["block"]="pprof.contentions.delay.001.pb.gz"
# mode2filename["threadcreate"]="pprof.contentions.delay.001.pb.gz"
# mode2filename["trace"]="pprof.contentions.delay.001.pb.gz"

modes=("goroutine" "heap" "profile")

mode=""
action=""
sec=60
port=8080
dirname="pprof"

while getopts "m:a:s:p:d:" opt; do
  case $opt in
  m)
    mode=$OPTARG
    ;;
  a)
    action=$OPTARG
    ;;
  s)
    sec=$OPTARG
    ;;
  p)
    port=$OPTARG
    ;;
  d)
    dirname=$OPTARG
    ;;
  \?)
    echo "Invalid option: -$OPTARG" >&2
    exit 1
    ;;
  esac
done

echo "mode=$mode, action=$action, sec=$sec, port=$port, dirname=$dirname"

listen() {
  mode=$1
  if [ "$mode" = "profile" ]; then
    suffix="seconds=$sec"
  elif [ "$mode" = "trace" ]; then
    suffix=""
  else
    suffix="debug=1"
  fi
  # go tool pprof "http://127.0.0.1:$pprof_port/debug/pprof/$mode?$suffix"
  bash -c "export PPROF_TMPDIR=\"/home/duk/code/RedisShake/demo/bench/$dirname\"; go tool pprof \"http://127.0.0.1:$pprof_port/debug/pprof/$mode?$suffix\"" &
}

if [ "$action" = "listen" ]; then
  # export PPROF_TMPDIR="/home/duk/code/RedisShake/demo/bench/$dirname"
  if [ "$mode" = "all" ]; then
    for m in ${modes[@]}; do
      
      listen $m
    done
  else
    listen $mode
  fi
elif [ "$action" = "web" ]; then
  filename=${mode2filename[$mode]}
  echo $filename
  go tool pprof -http=:$port "$logdir/$dirname/$filename"
fi
