2024-09-23 22:40:29 INF read_count=[2539522], read_ops=[39535.50], write_count=[2539522], write_ops=[39535.50], syncing aof, diff=[49162]
2024-09-23 22:40:34 INF read_count=[2868048], read_ops=[67147.59], write_count=[2868048], write_ops=[67147.59], syncing aof, diff=[65536]
2024-09-23 22:40:39 INF read_count=[3252994], read_ops=[76623.60], write_count=[3252994], write_ops=[76623.60], syncing aof, diff=[49152]
2024-09-23 22:40:44 INF read_count=[3639033], read_ops=[76632.83], write_count=[3639033], write_ops=[76632.83], syncing aof, diff=[65544]
2024-09-23 22:40:49 INF read_count=[4025071], read_ops=[76760.90], write_count=[4025071], write_ops=[76761.90], syncing aof, diff=[77825]
2024-09-23 22:40:54 INF read_count=[4413138], read_ops=[76543.50], write_count=[4413138], write_ops=[76543.50], syncing aof, diff=[81927]
2024-09-23 22:40:59 INF read_count=[4800823], read_ops=[77519.23], write_count=[4800823], write_ops=[77519.23], syncing aof, diff=[-16384]
2024-09-23 22:41:04 INF read_count=[5187711], read_ops=[76355.83], write_count=[5187711], write_ops=[76355.83], syncing aof, diff=[65540]
2024-09-23 22:41:09 INF read_count=[5575407], read_ops=[77493.36], write_count=[5575406], write_ops=[77493.36], syncing aof, diff=[4096]
2024-09-23 22:41:14 INF read_count=[5962596], read_ops=[75816.95], write_count=[5962596], write_ops=[75816.95], syncing aof, diff=[9]
2024-09-23 22:41:19 INF read_count=[6350038], read_ops=[77795.21], write_count=[6350038], write_ops=[77795.21], syncing aof, diff=[49152]
2024-09-23 22:41:24 INF read_count=[6736602], read_ops=[76660.02], write_count=[6736601], write_ops=[76659.02], syncing aof, diff=[20481]
2024-09-23 22:41:29 INF read_count=[7124455], read_ops=[78688.14], write_count=[7124455], write_ops=[78688.14], syncing aof, diff=[0]
2024-09-23 22:41:34 INF read_count=[7504720], read_ops=[70481.95], write_count=[7504720], write_ops=[70481.95], syncing aof, diff=[49165]
2024-09-23 22:41:39 INF read_count=[7841606], read_ops=[77968.29], write_count=[7841606], write_ops=[77968.29], syncing aof, diff=[16409]
2024-09-23 22:41:44 INF read_count=[8226240], read_ops=[76526.64], write_count=[8226240], write_ops=[76526.64], syncing aof, diff=[16399]
2024-09-23 22:41:49 INF read_count=[8613058], read_ops=[78654.33], write_count=[8613058], write_ops=[78654.33], syncing aof, diff=[-16384]
2024-09-23 22:41:54 INF read_count=[8997536], read_ops=[76448.92], write_count=[8997536], write_ops=[76448.92], syncing aof, diff=[32768]
2024-09-23 22:41:57 ERR write tcp 127.0.0.1:34096->127.0.0.1:38004: write: connection reset by peer
                        RedisShake/internal/client/redis.go:191 -> (*Redis).flush()
                        RedisShake/internal/client/redis.go:177 -> (*Redis).Send()
                        RedisShake/internal/reader/sync_standalone_reader.go:394 -> (*syncStandaloneReader).sendReplconfAck()
                        runtime/asm_amd64.s:1695 -> goexit()