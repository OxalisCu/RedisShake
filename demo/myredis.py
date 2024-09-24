import time
import redis
from util import Timer, get_free_port, PATH_REDIS_SERVER
from mylauncher import Launcher


class Redis:
    def __init__(self, args=None):
        if args is None:
            args = []
        self.host = "127.0.0.1"
        self.port = get_free_port()
        self.dirpath = f"/home/duk/code/RedisShake/demo/data/redis_{self.port}"
        args.extend(["--port", str(self.port)])
        args = [PATH_REDIS_SERVER] + args
        self.server = Launcher(args, work_dir=self.dirpath)
        self._wait_start()
        self.client = redis.Redis(host=self.host, port=self.port)
        self.client.config_set('notify-keyspace-events', 'AKEx')

    def __del__(self):
        if self.server.started:
            try:
                self.server.stop()
            except Exception as e:
                print("stop redis server failed")

    def _wait_start(self, timeout=5):
        timer = Timer()
        while True:
            try:
                r = redis.Redis(host=self.host, port=self.port)
                r.ping()
                return
            except redis.exceptions.ConnectionError:
                time.sleep(0.1)
            if timer.elapsed() > timeout:
                stderr = f"{self.dirpath}/stderr"
                with open(stderr, "rb") as f:
                    for line in f.readlines():
                        print(line.strip())
                raise TimeoutError("redis server not started")

    def do(self, *args):
        try:
            ret = self.client.execute_command(*args)
        except redis.exceptions.ResponseError as e:
            return f"-{str(e)}"
        return ret

    def pipeline(self):
        return self.client.pipeline(transaction=False)

    def get_address(self):
        return f"{self.host}:{self.port}"

    def is_cluster(self):
        return self.client.info()["cluster_enabled"]

    def dbsize(self):
        return self.client.dbsize()
