import json
from functools import wraps

from flask import Flask, jsonify, request, url_for, abort, Response

app = Flask(__name__)

def start():
    app.run()

@app.route('/health', methods=['GET'])
def healthcheck():
    response = jsonify({"message": "ok"})
    # ステータスコードは Created (201)
    response.status_code = 200
    return response