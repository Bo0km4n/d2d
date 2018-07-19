import os, sys
sys.path.append(os.getcwd())
import mysqlconn

sql_statement = "SELECT * FROM aggregated_log"

class AggregatedLog:
    def __init__(self):
        self.conn = mysqlconn.mysqlconn.get_mysql_conn()

    def fetch_all(self):
        cur = self.conn.cursor(dictionary=True)
        cur.execute(sql_statement)
        return cur.fetchall()

    def fetch_label(self, uid):
        cur = self.conn.cursor(dictionary=True)
        cur.execute("SELECT * FROM label WHERE uid = "+str(uid))
        return cur.fetchall()[0]['label']