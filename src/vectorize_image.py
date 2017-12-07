from keras.applications import VGG16
from keras.applications.vgg16 import preprocess_input
from keras.preprocessing import image
import numpy as np
import sys
from glob import glob
import scipy.sparse as sp
from scipy import io

def save_sparse_matrix(filename, x):
    x_coo = x.tocoo()
    row = x_coo.row
    col = x_coo.col
    data = x_coo.data
    shape = x_coo.shape
    np.savez(filename, row=row, col=col, data=data, shape=shape)

def load_sparse_matrix(filename):
    y = np.load(filename)
    z = sp.coo_matrix((y['data'], (y['row'], y['col'])), shape=y['shape'])
    return z

def save(arr, filename):
    with open(filename, 'w') as fd:
        fd.write(','.join(arr))

def vectorize(path, model):
    img = kimage.load_img(path, target_size=(224, 224))
    x = kimage.img_to_array(img)
    x = np.expand_dims(x, axis=0)
    x = preprocess_input(x)
    pred = model.predict(x)
    return pred.ravel()

def vectorize_all(root, model, batch_size=512):
    print("Will vectorize")
    files = glob(root)
    save(files, '../submission/images_order.csv')
    min_idx = 0
    max_idx = min_idx + batch_size
    total_max = len(files)
    px = 224
    n_dims = 512
    preds = sp.lil_matrix((len(files), n_dims))

    print("Total: {}".format(len(files)))
    while min_idx < total_max - 1:
        print(min_idx)
        X = np.zeros(((max_idx - min_idx), px, px, 3))
        # For each file in batch, 
        # load as row into X
        for i in range(min_idx, max_idx):
            file = files[i]
            try:
                img = image.load_img(file, target_size=(px, px))
                img_array = image.img_to_array(img)
                X[i - min_idx, :, :, :] = img_array
            except Exception as e:
                print(e)
        max_idx = i
        X = preprocess_input(X)
        these_preds = model.predict(X)
        shp = ((max_idx - min_idx) + 1, n_dims)
        preds[min_idx:max_idx + 1, :] = these_preds.reshape(shp)
        min_idx = max_idx
        max_idx = np.min((max_idx + batch_size, total_max))
    return preds


def main():
    model = VGG16(include_top=False, weights='imagenet', pooling='max')
    vecs = vectorize_all('../data/images/*.jpg', model, 512)
    save_sparse_matrix("../submission/images_vec.npz", vecs)

if __name__ == '__main__':
    main()