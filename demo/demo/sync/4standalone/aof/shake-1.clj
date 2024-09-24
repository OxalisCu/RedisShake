2024-09-23 22:40:32 INF read_count=[2735764], read_ops=[61778.91], write_count=[2735764], write_ops=[61778.91], syncing aof, diff=[16384]
2024-09-23 22:40:37 INF read_count=[3113844], read_ops=[77936.99], write_count=[3113844], write_ops=[77936.99], syncing aof, diff=[0]
2024-09-23 22:40:42 INF read_count=[3498790], read_ops=[78023.74], write_count=[3498790], write_ops=[78023.74], syncing aof, diff=[65536]
2024-09-23 22:40:47 INF read_count=[3884673], read_ops=[76401.38], write_count=[3884673], write_ops=[76401.38], syncing aof, diff=[49153]
2024-09-23 22:40:52 INF read_count=[4272583], read_ops=[77944.35], write_count=[4272583], write_ops=[77944.35], syncing aof, diff=[49152]
2024-09-23 22:40:57 INF read_count=[4659714], read_ops=[77399.86], write_count=[4659714], write_ops=[77398.86], syncing aof, diff=[32778]
2024-09-23 22:41:02 INF read_count=[5046844], read_ops=[78496.00], write_count=[5046844], write_ops=[78496.00], syncing aof, diff=[32768]
2024-09-23 22:41:07 INF read_count=[5433663], read_ops=[76615.16], write_count=[5433663], write_ops=[76615.16], syncing aof, diff=[32773]
2024-09-23 22:41:12 INF read_count=[5820793], read_ops=[78529.71], write_count=[5820793], write_ops=[78529.71], syncing aof, diff=[32779]
2024-09-23 22:41:17 INF read_count=[6208704], read_ops=[76614.24], write_count=[6208704], write_ops=[76614.24], syncing aof, diff=[14]
2024-09-23 22:41:22 INF read_count=[6596146], read_ops=[79178.77], write_count=[6596146], write_ops=[79178.77], syncing aof, diff=[61440]
2024-09-23 22:41:27 INF read_count=[6983104], read_ops=[78384.86], write_count=[6983103], write_ops=[78383.86], syncing aof, diff=[-12278]
2024-09-23 22:41:32 INF read_count=[7371343], read_ops=[78552.06], write_count=[7371343], write_ops=[78552.06], syncing aof, diff=[0]
2024-09-23 22:41:37 INF read_count=[7706201], read_ops=[63331.74], write_count=[7706201], write_ops=[63331.74], syncing aof, diff=[0]
2024-09-23 22:41:42 INF read_count=[8086623], read_ops=[75968.80], write_count=[8086623], write_ops=[75968.80], syncing aof, diff=[-16374]
2024-09-23 22:41:47 INF read_count=[8472036], read_ops=[77705.31], write_count=[8472036], write_ops=[77705.31], syncing aof, diff=[16396]
2024-09-23 22:41:52 INF read_count=[8858386], read_ops=[76899.94], write_count=[8858386], write_ops=[76899.94], syncing aof, diff=[32768]
2024-09-23 22:41:57 ERR write tcp 127.0.0.1:55212->127.0.0.1:38002: write: connection reset by peer
                        RedisShake/internal/client/redis.go:191 -> (*Redis).flush()
                        RedisShake/internal/client/redis.go:177 -> (*Redis).Send()
                        RedisShake/internal/reader/sync_standalone_reader.go:394 -> (*syncStandaloneReader).sendReplconfAck()
                        runtime/asm_amd64.s:1695 -> goexit()