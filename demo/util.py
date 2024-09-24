import random
import socket
import threading
import time
import toml
import os
import shutil
from enum import Enum

BENCH_TYPE_PATH = "/home/duk/code/RedisShake/demo/config/launchtype.txt"


class BenchType(Enum):
    sync_c2c = "sync-c2c"
    sync_c2s = "sync-c2s"
    sync_s2c = "sync-s2c"
    sync_s2s = "sync-s2s"
    scan_c2c = "scan-c2c"
    scan_c2s = "scan-c2s"
    scan_s2c = "scan-s2c"
    scan_s2s = "scan-s2s"


def src_cluster(t: BenchType):
    if t in (
        BenchType.sync_c2c,
        BenchType.sync_c2s,
        BenchType.scan_c2c,
        BenchType.scan_c2s,
    ):
        return True
    return False


def dst_cluster(t: BenchType):
    if t in (
        BenchType.sync_c2c,
        BenchType.sync_s2c,
        BenchType.scan_c2c,
        BenchType.scan_s2c,
    ):
        return True
    return False


def is_sync(t: BenchType):
    if t in (
        BenchType.sync_c2c,
        BenchType.sync_c2s,
        BenchType.sync_s2c,
        BenchType.sync_s2s,
    ):
        return True
    return False


def get_bench_type(t: str):
    for bench_type in BenchType:
        if bench_type.value in t:
            return bench_type


def get_t_from_file():
    with open(BENCH_TYPE_PATH, "r", encoding="utf-8") as f:
        bench_type = f.read()
    return get_bench_type(bench_type)


class Timer:
    def __init__(self):
        self.start_time = time.perf_counter()

    def elapsed(self):
        return time.perf_counter() - self.start_time


def is_port_available(port: int) -> bool:
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        try:
            s.bind(("localhost", port))
            s.close()
            return True
        except OSError:
            return False


MIN_PORT = 20000
MAX_PORT = 40000

port_cursor = random.choice(range(MIN_PORT, MAX_PORT, 1000))

g_lock = threading.Lock()


def get_free_port():
    global port_cursor
    global g_lock
    with g_lock:
        while True:
            port_cursor += 1
            if port_cursor == MAX_PORT:
                port_cursor = MIN_PORT

            if is_port_available(port_cursor):
                return port_cursor


def create_sync_opts(src, dst):
    d = {
        "type": "sync",
        "sync_reader": {
            "cluster": src.is_cluster(),
            "address": src.get_address(),
            "sync_rdb": True,
            "sync_aof": True,
        },
        "redis_writer": {
            "cluster": dst.is_cluster(),
            "address": dst.get_address(),
        },
    }
    return d


def create_scan_opts(src, dst):
    d = {
        "type": "scan",
        "scan_reader": {
            "cluster": src.is_cluster(),
            "address": src.get_address(),
            "scan": True,
            "ksn": True,
        },
        "redis_writer": {
            "cluster": dst.is_cluster(),
            "address": dst.get_address(),
        },
    }
    return d

CONFIG_PATH = "/home/duk/code/RedisShake/demo"
CONFIG_DIRNAME = "config"

def save_toml_file(opts, filename):
    pprof_port = 40000
    opts["advanced"] = {"pprof_port": pprof_port, "log_level": "info"}

    dirpath = f"/home/duk/code/RedisShake/demo/{CONFIG_DIRNAME}"
    filepath = f"{dirpath}/{filename}.toml"
    with open(filepath, "w", encoding="utf-8") as f:
        toml.dump(opts, f)
    return filepath


PATH_REDIS_SERVER = shutil.which("redis-server")
