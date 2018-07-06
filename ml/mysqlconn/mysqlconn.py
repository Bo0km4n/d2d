from urllib.parse import urlparse
import mysql.connector

def get_mysql_conn():
    url = urlparse('mysql://root:password@localhost:13306/d2d')

    conn = mysql.connector.connect(
        host = url.hostname or 'localhost',
        port = url.port or 13306,
        user = url.username or 'root',
        password = url.password or '',
        database = url.path[1:],
    )
    return conn
