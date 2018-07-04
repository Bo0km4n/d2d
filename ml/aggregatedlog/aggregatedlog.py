import os, sys
sys.path.append(os.getcwd())
import mysqlconn
class AggregatedLog:
    def __init__(self):
        self.conn = mysqlconn.mysqlconn.get_mysql_conn()

    def fetch_all(self):
        cur = self.conn.cursor()
        cur.execute("SELECT * FROM aggregated_log")
        return cur.fetchall()
