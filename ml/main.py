import os, sys
sys.path.append(os.getcwd())
import aggregatedlog
import batch
logger = aggregatedlog.aggregatedlog.AggregatedLog()

logs_data = logger.fetch_all()

x_train, y_train = batch.formatter.format_data(logs_data)

svm = batch.svm.SVM(x_train, y_train)
svm.exec_fit()
svm.accuracy_pred()