import os, sys
sys.path.append(os.getcwd())
import aggregatedlog
import batch
import argparse
import api
parser = argparse.ArgumentParser()
parser.add_argument('-m', '--mode', type=str)
args = parser.parse_args()

if args.mode == 'api':
    api.app.start()

elif args.mode == 'svm':
    logger = aggregatedlog.aggregatedlog.AggregatedLog()
    logs_data = logger.fetch_all()
    x_train, y_train = batch.formatter.format_data(logs_data)
    svm = batch.svm.SVM(x_train, y_train)
    svm.exec_fit()
    svm.accuracy_pred()
    svm.save_model()