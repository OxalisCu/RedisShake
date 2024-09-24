import time
import redis
from myredis import Redis

slot_num = 16384


class Cluster:
    def __init__(self, num):
        self.num = num
        self.nodes = []
        slots_per_node = slot_num // num
        for i in range(num):
            r = Redis(args=["--cluster-enabled", "yes"])
            self.nodes.append(r)
            start_slot = i * slots_per_node
            end_slot = start_slot + slots_per_node - 1 if i != num - 1 else slot_num - 1
            ret = self.nodes[i].do(
                "cluster", "addslots", *range(start_slot, end_slot + 1)
            )

        for i in range(num):
            for j in range(i + 1, num):
                ret = self.nodes[i].do("cluster", "meet", self.nodes[j].host, self.nodes[j].port)

        startup_nodes = [
            redis.cluster.ClusterNode(node.host, node.port) for node in self.nodes
        ]
        self.client = redis.RedisCluster(
            startup_nodes=startup_nodes,
            require_full_coverage=True,
        )

        for _ in range(20):
            cluster_info = self.client.cluster_info()
            if cluster_info["cluster_state"] == "ok":
                break
            time.sleep(0.1)
        else:
            raise Exception("Cluster initialization failed.")
        print(f"cluster state: {self.client.cluster_info()['cluster_state']}")

    def do(self, *args):
        try:
            ret = self.client.execute_command(*args)
        except redis.exceptions.ResponseError as e:
            return f"-{str(e)}"
        return ret

    def pipeline(self):
        return self.client.pipeline(transaction=False)

    def get_address(self):
        return self.nodes[0].get_address()

    def dbsize(self):
        size = 0
        for node in self.nodes:
            size += node.dbsize()
        return size

    @staticmethod
    def is_cluster():
        return True
