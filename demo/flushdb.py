import redis
from enum import Enum
import config
from util import BenchType


class FlushType(Enum):
    src = "src"
    dst = "dst"
    both = "both"


def delete_cluster_data(host, port):
    master_node = redis.Redis(host=host, port=port)
    cluster_nodes = master_node.execute_command("CLUSTER NODES")
    cluster_addresses = cluster_nodes.keys()

    for node_address in cluster_addresses:
        node_host, node_port = node_address.split(":")
        delete_redis_data(node_host, node_port)


def delete_redis_data(host, port):
    node = redis.Redis(host=host, port=port)
    node.flushall()


def delete_data(cluster, host, port):
    if cluster:
        delete_cluster_data(host, port)
    else:
        delete_redis_data(host, port)


def flushdb(flush_type: FlushType):
    res = config.get_config()
    if flush_type == FlushType.src:
        delete_data(res["src"]["cluster"], res["src"]["host"], res["src"]["port"])
    elif flush_type == FlushType.dst:
        delete_data(res["dst"]["cluster"], res["dst"]["host"], res["dst"]["port"])
    else:
        delete_data(res["src"]["cluster"], res["src"]["host"], res["src"]["port"])
        delete_data(res["dst"]["cluster"], res["dst"]["host"], res["dst"]["port"])
    print("flushed data")


flush_type = FlushType.both
flushdb(flush_type)
