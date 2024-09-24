from mycluster import Cluster
from myredis import Redis
import argparse
from util import (
    BenchType,
    BENCH_TYPE_PATH,
    src_cluster,
    dst_cluster,
    is_sync,
    create_sync_opts,
    create_scan_opts,
    save_toml_file,
    get_bench_type,
)


def launch_redis_instances(t: BenchType, src_num: int, dst_num: int):
    print("launching src")
    src = Cluster(src_num) if src_cluster(t) else Redis()
    print("launching dst")
    dst = Cluster(dst_num) if dst_cluster(t) else Redis()
    opts = create_sync_opts(src, dst) if is_sync(t) else create_scan_opts(src, dst)
    filepath = save_toml_file(opts, t.value)

    try:
        while True:
            pass
    except KeyboardInterrupt:
        pass


def get_t_and_save(filepath: str) -> BenchType:
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "-type",
        required=True,
        help="type of benchmark to run",
    )

    bench_type = parser.parse_args().type
    t = get_bench_type(bench_type)
    

    with open(filepath, "w", encoding="utf-8") as file:
        file.write(t.value)

    return t


t = get_t_and_save(BENCH_TYPE_PATH)
launch_redis_instances(t, 4, 4)
