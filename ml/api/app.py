import os, sys
sys.path.append(os.getcwd())

import batch.svm
import batch.formatter
import json
import numpy as np
from functools import wraps
from flask import Flask, jsonify, request, url_for, abort, Response

app = Flask(__name__)
clf = batch.svm.SVM()

def consumes(content_type):
    def _consumes(function):
        @wraps(function)
        def __consumes(*argv, **keywords):
            if request.headers['Content-Type'] != content_type:
                abort(400)
            return function(*argv, **keywords)
        return __consumes
    return _consumes

def start():
    app.run()

@app.route('/health', methods=['GET'])
def healthcheck():
    response = jsonify({"message": "ok"})
    # ステータスコードは Created (201)
    response.status_code = 200
    return response

@app.route('/api/v1/predict', methods=['POST'])
@consumes('application/json')
def predict():
    content_body = json.loads(request.data)
    clf.load_model()

    # convert json object to array
    b = batch.formatter.extract_tag_data(content_body)
    b_array = batch.formatter.min_max(np.array(b))
    
    b_array = np.array([b_array])
    result = clf.pred(b_array)

    resp_body = {"predict": int(result[0])}
    print(resp_body)
    response = jsonify(resp_body)
    response.status_code = 201
    return response

def get_label(label):
    return label