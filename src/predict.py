# -*- coding: utf-8 -*-
from urllib import request
from uuid import uuid4

import numpy as np
from keras.applications import VGG19
from keras.applications.vgg19 import preprocess_input
from keras.engine import Model
from keras.preprocessing import image
from sklearn.neighbors import NearestNeighbors

from settings import config
from sparce import load_sparse_matrix


def _vectorize(path, model):
    img = image.load_img(path, target_size=(224, 224))
    x = image.img_to_array(img)
    x = np.expand_dims(x, axis=0)
    x = preprocess_input(x)
    pred = model.predict(x)
    return pred.ravel()


def _similar(vec, knn, filenames, n_neighbors=6):
    dist, indices = knn.kneighbors(vec.reshape(1, -1), n_neighbors=n_neighbors)
    dist, indices = dist.flatten(), indices.flatten()
    return [(filenames[indices[i]], dist[i]) for i in range(len(indices))]


def load_predictor(owner_id):
    filenames = open(config.images_order(owner_id), 'r').readline().split(',')
    vecs = load_sparse_matrix(config.vectors_path(owner_id))
    base_model = VGG19(weights='imagenet')
    # Read about fc1 here http://cs231n.github.io/convolutional-networks/
    model = Model(inputs=base_model.input, outputs=base_model.get_layer('fc1').output)
    knn = NearestNeighbors(metric='cosine', algorithm='brute')
    knn.fit(vecs)

    def similarity(file_path, n_neighbors=6):
        vec = _vectorize(file_path, model)
        return _similar(vec, knn, filenames, n_neighbors)

    return similarity


def random(owner_id):
    filenames = open(config.images_order(owner_id), 'r').readline().split(',')

    def rf():
        return np.random.choice(filenames)

    return rf


def download(url):
    f = "/tmp/images/{}".format(str(uuid4()))
    request.urlretrieve(url, f)
    return f
