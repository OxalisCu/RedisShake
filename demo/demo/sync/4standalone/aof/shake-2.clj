2024-09-23 22:40:33 INF read_count=[2767309], read_ops=[62051.91], write_count=[2767309], write_ops=[62051.91], syncing aof, diff=[8]
2024-09-23 22:40:38 INF read_count=[3151943], read_ops=[76744.19], write_count=[3151943], write_ops=[76744.19], syncing aof, diff=[16403]
2024-09-23 22:40:43 INF read_count=[3536265], read_ops=[77134.93], write_count=[3536265], write_ops=[77134.93], syncing aof, diff=[49152]
2024-09-23 22:40:48 INF read_count=[3923239], read_ops=[76928.39], write_count=[3923239], write_ops=[76928.39], syncing aof, diff=[32784]
2024-09-23 22:40:53 INF read_count=[4310526], read_ops=[78188.37], write_count=[4310526], write_ops=[78188.37], syncing aof, diff=[3]
2024-09-23 22:40:58 INF read_count=[4697851], read_ops=[76871.61], write_count=[4697850], write_ops=[76871.61], syncing aof, diff=[32768]
2024-09-23 22:41:03 INF read_count=[5083850], read_ops=[77179.29], write_count=[5083850], write_ops=[77179.29], syncing aof, diff=[32768]
2024-09-23 22:41:08 INF read_count=[5472541], read_ops=[76770.05], write_count=[5472541], write_ops=[76770.05], syncing aof, diff=[-16384]
2024-09-23 22:41:13 INF read_count=[5858267], read_ops=[78062.96], write_count=[5858267], write_ops=[78062.96], syncing aof, diff=[16384]
2024-09-23 22:41:18 INF read_count=[6246802], read_ops=[76915.19], write_count=[6246802], write_ops=[76915.19], syncing aof, diff=[65553]
2024-09-23 22:41:23 INF read_count=[6633620], read_ops=[77086.91], write_count=[6633620], write_ops=[77086.91], syncing aof, diff=[16384]
2024-09-23 22:41:28 INF read_count=[7020974], read_ops=[77596.71], write_count=[7020974], write_ops=[77596.71], syncing aof, diff=[-4081]
2024-09-23 22:41:33 INF read_count=[7408662], read_ops=[78209.38], write_count=[7408662], write_ops=[78209.38], syncing aof, diff=[49153]
2024-09-23 22:41:38 INF read_count=[7738526], read_ops=[62869.52], write_count=[7738526], write_ops=[62869.52], syncing aof, diff=[23]
2024-09-23 22:41:43 INF read_count=[8123628], read_ops=[77206.42], write_count=[8123628], write_ops=[77206.42], syncing aof, diff=[65547]
2024-09-23 22:41:48 INF read_count=[8508886], read_ops=[77088.33], write_count=[8508886], write_ops=[77088.33], syncing aof, diff=[16401]
2024-09-23 22:41:53 INF read_count=[8895860], read_ops=[78168.55], write_count=[8895860], write_ops=[78168.55], syncing aof, diff=[32768]
2024-09-23 22:41:57 ERR write tcp 127.0.0.1:60470->127.0.0.1:38001: write: connection reset by peer
                        RedisShake/internal/client/redis.go:191 -> (*Redis).flush()
                        RedisShake/internal/client/redis.go:177 -> (*Redis).Send()
                        RedisShake/internal/reader/sync_standalone_reader.go:394 -> (*syncStandaloneReader).sendReplconfAck()
                        runtime/asm_amd64.s:1695 -> goexit()