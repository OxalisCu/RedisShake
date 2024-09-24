#!/bin/bash

datadir=/home/duk/code/RedisShake/demo/data/
shakedir=/home/duk/code/RedisShake/cmd/redis-shake/data

if [ "$1" = "clean" ]; then
  rm -rf /home/duk/code/RedisShake/data
  rm -rf /home/duk/code/RedisShake/demo/data/*
  rm -rf /home/duk/code/RedisShake/cmd/redis-shake/data/*
elif [ "$1" = "kill" ]; then
  killall -9 redis-server
else
  echo -e "please select clean mode: clean, kill\n"
fi
