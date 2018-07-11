import os, sys
sys.path.append(os.getcwd())
from sklearn.svm import SVC
from sklearn.metrics import accuracy_score

class SVM:
    def __init__(self, x_train, y_train):
        self.x_train = x_train
        self.y_train = y_train
        self.model = SVC(kernel='linear', random_state=None)
    
    def exec_fit(self):
        self.model.fit(self.x_train, self.y_train)
    
    def accuracy_pred(self):
        pred_train = self.model.predict(self.x_train)
        accuracy_train = accuracy_score(self.y_train, pred_train)
        print('トレーニングデータに対する正解率： %.2f' % accuracy_train)
    