from glob import glob

import numpy as np
import scipy.sparse as sp
from keras.applications import VGG19
from keras.applications.vgg19 import preprocess_input
from keras.engine import Model
from keras.preprocessing import image

from settings import config
from sparce import save_sparse_matrix


def save(arr, filename):
    with open(filename, 'w') as fd:
        fd.write(','.join(arr))


def vectorize_all(files, model, px=224, n_dims=512, batch_size=512):
    print("Will vectorize")
    min_idx = 0
    max_idx = min_idx + batch_size
    total_max = len(files)
    preds = sp.lil_matrix((len(files), n_dims))

    print("Total: {}".format(len(files)))
    while min_idx < total_max - 1:
        print(min_idx)
        X = np.zeros(((max_idx - min_idx), px, px, 3))
        # For each file in batch, 
        # load as row into X
        i = 0
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


def main(dir_name):
    files = glob(config.images_glob_path(dir_name))
    save(files, config.images_order(dir_name))

    base_model = VGG19(weights='imagenet')
    # Read about fc1 here http://cs231n.github.io/convolutional-networks/
    model = Model(inputs=base_model.input, outputs=base_model.get_layer('fc1').output)

    vecs = vectorize_all(files, model, n_dims=4096)
    save_sparse_matrix(config.vectors_path(dir_name), vecs)
    print('Finished')
