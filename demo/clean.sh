#!/bin/bash

# srcdir read from .env

if [ "$1" = "clean" ]; then
  rm -rf $srcdir/data*
  rm -rf $srcdir/demo/data/*
  rm -rf $srcdir/cmd/redis-shake/data/*
elif [ "$1" = "kill" ]; then
  killall -9 redis-server
else
  echo -e "please select clean mode: clean, kill\n"
fi
