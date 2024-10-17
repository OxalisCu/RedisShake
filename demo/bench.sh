taskset -c 0-15 redis-benchmark \
  --threads 16 -r 10000000 -n 1000000000 -t set \
  -h r-bp16uk6iq1sfxg6zbn.redis.rds.aliyuncs.com \
  -a 'username:D3xP!4@AS!d@6q' \
  --cluster -c 256 -d 8 -P 2

# taskset -c 0-7 redis-benchmark \
#   --threads 8 -r 10000000 -n 1000000000 -t set \
#   -h r-bp16uk6iq1sfxg6zbn.redis.rds.aliyuncs.com \
#   -a 'username:D3xP!4@AS!d@6q' \
#   --cluster -c 256 -d 8 -P 2