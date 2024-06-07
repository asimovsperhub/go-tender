import json, os
import logging
import subprocess
import uuid
from logging.handlers import RotatingFileHandler
import pymysql
import redis
from concurrent.futures import ThreadPoolExecutor
import sys

from minio import Minio

sys.path.insert(
    0,
    os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
)
host = "127.0.0.1"
cache_cli = redis.StrictRedis(  # 缓存客户端对象
    host=host,
    db=1,
    password='asimov@77',
    socket_connect_timeout=1  # 链接超时设置为1秒
)


# class DecimalEncoder(json.JSONEncoder):
#     def default(self, o):
#         if isinstance(o, decimal.Decimal):
#             return float(o)
#         super(DecimalEncoder, self).default(o)


class Db(object):

    def __init__(self, database: str = 'tender'):
        try:
            self.db = pymysql.connect(
                # 42.193.247.183
                host=host,
                port=3306,
                user="root",
                password="JyMysql@007", database=database, charset="utf8"
            )
        except Exception as e:
            logger.error(e)
            os._exit(0)

    def search(self, sql: str):
        cs = self.db.cursor()
        try:
            cs.execute(sql)
            r = cs.execute(sql)
            logger.info(r)
            return r
        except Exception as e:
            logger.error(f"{sql} err: {e}")
            return ()
        finally:
            cs.close()

    def insert(self, sql: str):
        cs = self.db.cursor()
        try:
            cs.execute(sql)
            self.db.commit()
            return cs.lastrowid
        except Exception as e:
            print(e)
            self.db.rollback()
        finally:
            cs.close()

    def test_conn(self,database: str = 'tender'):
        try:
            logger.info("ping-------------> mysql")
            self.db.ping()
        except:
            try:
                logger.info("重新连接mysql---------------->")
                self.db = pymysql.connect(
                    # 42.193.247.183
                    host=host,
                    port=3306,
                    user="root",
                    password="JyMysql@007", database=database, charset="utf8"
                )
            except Exception as e:
                logger.error(e)
                os._exit(0)
    def close_cs_db(self):
        self.db.close()


class XLogger(logging.Logger):
    """
    自定义日志对象, 快速的实现文件输出, 以及捕获异常堆栈
    """

    def __init__(
            self, name=None,
            filepath=None,
            level=logging.DEBUG,
            fmt="%(asctime)s %(levelname)0.4s %(filename)s:%(lineno)d %(message)s",
            date_fmt="%Y-%m-%d %H:%M:%S",
            max_bytes=1024 * 1024 * 10,
            backup_count=10
    ):
        super().__init__(name)
        formatter = logging.Formatter(fmt=fmt, datefmt=date_fmt)
        stream_handler = logging.StreamHandler()
        stream_handler.setFormatter(formatter)
        self.addHandler(stream_handler)
        if filepath:
            file_handler = RotatingFileHandler(filename=filepath, maxBytes=max_bytes, backupCount=backup_count,
                                               encoding='utf-8')
            file_handler.setFormatter(formatter)
            self.addHandler(file_handler)
        self.setLevel(level=level)

    def error(self, msg, *args, **kwargs):
        """
        Log 'msg % args' with severity 'ERROR'.

        To pass exception information, use the keyword argument exc_info with
        a true value, e.g.

        logger.error("Houston, we have a %s", "major problem", exc_info=1)
        """
        kwargs.setdefault('exc_info', True)
        if self.isEnabledFor(logging.ERROR):
            self._log(logging.ERROR, msg, args, **kwargs)


def S3Upload(bucket: str, content, filename: str, length: int = -1):
    useSSL = False
    client = Minio(
        f"{host}:9000",
        access_key="JvTpb5Gnt4qCL69v",
        secret_key="AnyVSSQK2aTWM2QmE3HjZQB1040gYB5L", secure=useSSL
    )
    found = client.bucket_exists(bucket)
    if not found:
        client.make_bucket(bucket)
    # 存储桶，路径，[]byte,大小
    if length > 0:
        result = client.put_object(
            bucket, filename, content, length,
        )
    else:
        # Upload unknown sized data  length=-1, part_size=10*1024*1024, 大文件内存可能出问题
        result = client.put_object(
            bucket, filename, content, length=-1, part_size=10 * 1024 * 1024,
        )
    return result


db = Db()

class Cutting:
    def __init__(self):
        self.pool = ThreadPoolExecutor(max_workers=3)  # 线程池执行器

    def cutting(self, video: str, id: str):
        code = video.split(".")[-1]
        file_name = str(uuid.uuid1()) + "." + code
        short_video = "/tmp/" + file_name
        cmd = f"ffmpeg -ss 1 -i {video} -t 15 -vcodec h264 {short_video} -y"
        subprocess.call(cmd, shell=True)
        filesize = os.path.getsize(short_video)
        with open(short_video, 'rb') as f:
            result = S3Upload("tender", f, f"/{code}/" + file_name, length=int(filesize))
            print(result.object_name, result.etag, result.version_id)
            short_video_url = "https://biaoziku.com/tender/" + code + "/" + file_name
        img_name = str(uuid.uuid1()) + ".png"
        cv = "/tmp/" + img_name
        cmd1 = f"ffmpeg -i {video} -filter_complex '[0]select=gte(n\,1)[s0]' -map [s0] -f image2 -vcodec mjpeg -vframes 1 {cv} -y"
        subprocess.call(cmd1, shell=True)
        imgsize = os.path.getsize(cv)
        with open(cv, 'rb') as f:
            result = S3Upload("tender", f, "/PNG/" + img_name, length=int(imgsize))
            print(result.object_name, result.etag, result.version_id)
            img_url = "https://biaoziku.com/tender/" + "PNG" + "/" + img_name
        sql = f"update member_knowledge set shortvideo_url = '{short_video_url}',cover_url='{img_url}',cut_status=1 where id = {id};"
        print("sql---------->", sql)
        db.test_conn()
        db.insert(sql)

    def get_cutting(self):
        while True:
            try:
                video = cache_cli.blpop("cutting")
                video = video[1].decode("utf-8")
                logger.info(f"video :{video}")
                sp = video.split("_/tmp")
                id = sp[0]
                video = "/tmp" + sp[-1]
                logger.info(f"id:{id},video :{video}")
                self.cutting(video, id)
            except Exception as e:
                logger.error(e)

    def run(self):
        [self.pool.submit(self.get_cutting) for i in range(3)]  # bluechip
        # wait=True 会 join() 阻塞等待其他线程完成
        self.pool.shutdown(wait=True)


if __name__ == '__main__':
    PROJECT_DIR = os.path.dirname(os.path.dirname(__file__))
    logger = XLogger(
        name="cutting",
        filepath=os.path.join(PROJECT_DIR, "cutting.txt")
    )
    coin = Cutting()
    coin.run()
