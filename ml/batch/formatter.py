import os, sys
sys.path.append(os.getcwd())
from sklearn.preprocessing import StandardScaler 
import numpy as np

extract_tags = [
    'elapsed_time',
    'delta_accel_x',
    'delta_accel_y',
    'delta_accel_z',
    'delta_gyro_x',
    'delta_gyro_y',
    'delta_gyro_z',
    'delta_mag_x',
    'delta_mag_y',
    'delta_mag_z'
]

label_tag = 'label'

# return x_train, y_train
def format_data(data):
    x_train = []
    y_train = []

    for v in data:
        # TODO
        # データをx_trainとy_trainに分割する
        
        ok, label = extract_label(v)
        if ok:
            y_train.append(label)
            x_train.append(extract_tag_data(v))

    x_train_std = min_max(np.array(x_train))

    return x_train_std, y_train

def extract_tag_data(v):
    values = []
    for i, k in enumerate(v):
        try:
            if k == label_tag or k == 'uuid' or k == 'id':
                continue
            values.append(v[k])
        except KeyError:
            continue

    return values

def extract_label(v):
    try:
        if v[label_tag] == '':
            return False, ''
        else:
            return True, v[label_tag]
    except KeyError:
        return False, ''

def min_max(x, axis=None):
    min = x.min(axis=axis, keepdims=True)
    max = x.max(axis=axis, keepdims=True)
    result = (x-min)/(max-min)
    return result