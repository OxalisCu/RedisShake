from resp_benchmark import Benchmark
from config import get_config


def bench():
    res = get_config()
    bm = Benchmark(
        host=res["src"]["host"], port=res["src"]["port"], cluster=res["src"]["cluster"]
    )
    result = bm.bench(
        "SET {key sequence 10000000} {value 64}", count=100000000, connections=32
    )
    print(result.qps, result.avg_latency_ms, result.p99_latency_ms)


bench()
