from keras.models import Sequential
from keras.layers.convolutional import Conv2D
from keras.layers.pooling import MaxPool2D
from keras.optimizers import Adam

from keras.layers.core import Dense, Activation, Dropout, Flatten
from keras.utils import plot_model
from keras.callbacks import TensorBoard

from keras.datasets import cifar10
from keras.utils import np_utils

(X_train, y_train),(X_test, y_test) = cifar10.load_data()

nb_classes = 10
batch_size = 128
nb_epoch = 20

# floatに型変換
X_train = X_train.astype('float32')

X_test = X_test.astype('float32')
# 各画素値を正規化
X_train /= 255.0
X_test /= 255.

Y_train = np_utils.to_categorical(y_train, nb_classes)
print(Y_train)
Y_test = np_utils.to_categorical(y_test, nb_classes)

# モデルの定義
model = Sequential()

model.add(Conv2D(32,3,input_shape=(32,32,3)))
model.add(Activation('relu'))
model.add(Conv2D(32,3))
model.add(Activation('relu'))
model.add(MaxPool2D(pool_size=(2,2)))

model.add(Conv2D(64,3))
model.add(Activation('relu'))
model.add(MaxPool2D(pool_size=(2,2)))

model.add(Flatten())
model.add(Dense(1024))
model.add(Activation('relu'))
model.add(Dropout(1.0))

model.add(Dense(nb_classes, activation='softmax'))

adam = Adam(lr=1e-4)

model.compile(optimizer=adam, loss='categorical_crossentropy', metrics=["accuracy"])

history = model.fit(X_train, Y_train, batch_size=batch_size, nb_epoch=nb_epoch, verbose=1, validation_split=0.1)
# モデルをプロット
plot_model(model, to_file='./model2.png')

json_string = model.to_json()
open(os.path.join('./', 'cnn_model.json'), 'w').write(json_string)

model.save_weights(os.path.join('./', 'cnn_model_weight.hdf5'))