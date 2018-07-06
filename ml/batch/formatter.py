import os, sys
sys.path.append(os.getcwd())

# return x_train, y_train
def format_aggregated_data(data):
    x_train = []
    y_train = []

    for v in data:
        # TODO
        # データをx_trainとy_trainに分割する
        x_train.append([])