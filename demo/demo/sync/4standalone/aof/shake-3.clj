2024-09-23 22:40:33 INF read_count=[2814244], read_ops=[71932.97], write_count=[2814244], write_ops=[71932.97], syncing aof, diff=[10]
2024-09-23 22:40:38 INF read_count=[3198878], read_ops=[77098.00], write_count=[3198878], write_ops=[77098.00], syncing aof, diff=[21]
2024-09-23 22:40:43 INF read_count=[3584448], read_ops=[77254.50], write_count=[3584448], write_ops=[77254.50], syncing aof, diff=[-16363]
2024-09-23 22:40:48 INF read_count=[3970954], read_ops=[78080.19], write_count=[3970954], write_ops=[78080.19], syncing aof, diff=[0]
2024-09-23 22:40:53 INF read_count=[4357773], read_ops=[76783.98], write_count=[4357773], write_ops=[76783.98], syncing aof, diff=[13]
2024-09-23 22:40:58 INF read_count=[4745839], read_ops=[78159.01], write_count=[4745839], write_ops=[78159.01], syncing aof, diff=[32768]
2024-09-23 22:41:03 INF read_count=[5131097], read_ops=[76246.83], write_count=[5131097], write_ops=[76246.83], syncing aof, diff=[65536]
2024-09-23 22:41:08 INF read_count=[5519944], read_ops=[78312.63], write_count=[5519944], write_ops=[78312.63], syncing aof, diff=[49152]
2024-09-23 22:41:13 INF read_count=[5905670], read_ops=[76954.96], write_count=[5905670], write_ops=[76953.96], syncing aof, diff=[81920]
2024-09-23 22:41:18 INF read_count=[6295297], read_ops=[78788.24], write_count=[6295297], write_ops=[78788.24], syncing aof, diff=[32768]
2024-09-23 22:41:23 INF read_count=[6681804], read_ops=[76961.72], write_count=[6681804], write_ops=[76961.72], syncing aof, diff=[12]
2024-09-23 22:41:28 INF read_count=[7070026], read_ops=[78687.88], write_count=[7070026], write_ops=[78687.88], syncing aof, diff=[0]
2024-09-23 22:41:33 INF read_count=[7456533], read_ops=[76926.61], write_count=[7456533], write_ops=[76926.61], syncing aof, diff=[49175]
2024-09-23 22:41:38 INF read_count=[7785305], read_ops=[72228.13], write_count=[7785305], write_ops=[72228.13], syncing aof, diff=[73728]
2024-09-23 22:41:43 INF read_count=[8171343], read_ops=[76930.39], write_count=[8171343], write_ops=[76930.39], syncing aof, diff=[65536]
2024-09-23 22:41:48 INF read_count=[8556913], read_ops=[77088.37], write_count=[8556913], write_ops=[77088.37], syncing aof, diff=[0]
2024-09-23 22:41:53 INF read_count=[8942951], read_ops=[78544.41], write_count=[8942951], write_ops=[78543.41], syncing aof, diff=[65536]
2024-09-23 22:41:57 ERR write tcp 127.0.0.1:58904->127.0.0.1:38003: write: connection reset by peer
                        RedisShake/internal/client/redis.go:191 -> (*Redis).flush()
                        RedisShake/internal/client/redis.go:177 -> (*Redis).Send()
                        RedisShake/internal/reader/sync_standalone_reader.go:394 -> (*syncStandaloneReader).sendReplconfAck()
                        runtime/asm_amd64.s:1695 -> goexit()