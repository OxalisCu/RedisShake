import toml
from util import get_t_from_file, CONFIG_PATH, CONFIG_DIRNAME

def parse_address(address: str):
    info = str.split(address, ":")
    return info


def get_config():
    t = get_t_from_file()
    configpath = {
        "sync-c2c": {"type": "sync", "config_path": f"{CONFIG_PATH}/{CONFIG_DIRNAME}/sync-c2c.toml"},
        "scan-c2c": {"type": "scan", "config_path": f"{CONFIG_PATH}/{CONFIG_DIRNAME}/scan-c2c.toml"},
    }
    configs = toml.load(configpath[t.value]["config_path"])
    reader_address = configs[f"{configpath[t.value]['type']}_reader"]["address"]
    writer_address = configs["redis_writer"]["address"]
    reader_info = str.split(reader_address, ":")
    writer_info = str.split(writer_address, ":")
    reader_cluster = configs[f"{configpath[t.value]['type']}_reader"]["cluster"]
    writer_cluster = configs["redis_writer"]["cluster"]
    res = {
        "src": {
            "host": reader_info[0],
            "port": int(reader_info[1]),
            "cluster": reader_cluster,
        },
        "dst": {
            "host": writer_info[0],
            "port": int(writer_info[1]),
            "cluster": writer_cluster,
        },
        # "pprof_port": configs["advanced"]["pprof_port"],
        "launch_type": t,
    }
    return res
