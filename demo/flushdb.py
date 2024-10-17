import redis
from enum import Enum
import config
import argparse


class FlushType(Enum):
    src = "src"
    dst = "dst"
    both = "both"


def delete_cluster_data(redis_config):
    master_node = redis.Redis(host=redis_config["host"], port=redis_config["port"], username=redis_config["username"], password=redis_config["password"])
    cluster_nodes = master_node.execute_command("CLUSTER NODES")
    cluster_addresses = cluster_nodes.keys()

    for node_address in cluster_addresses:
        node_host, node_port = node_address.split(":")
        node_config = {
            "host": node_host,
            "port": node_port,
            "username": redis_config["username"],
            "password": redis_config["password"],
        }
        delete_redis_data(node_config)


def delete_redis_data(redis_config):
    node = redis.Redis(host=redis_config["host"], port=redis_config["port"], username=redis_config["username"], password=redis_config["password"])
    node.flushall()


def delete_data(redis_config):
    if redis_config["cluster"]:
        delete_cluster_data(redis_config)
    else:
        delete_redis_data(redis_config)


def flushdb(flush_type: FlushType):
    res = config.get_config()
    if flush_type == FlushType.src:
        delete_data(res["src"])
    elif flush_type == FlushType.dst:
        delete_data(res["dst"])
    else:
        delete_data(res["src"])
        delete_data(res["dst"])
    print("flushed data")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Flush data from redis")
    parser.add_argument("-t", type=FlushType, choices=list(FlushType), required=True)
    args = parser.parse_args()
    flushdb(args.t)
