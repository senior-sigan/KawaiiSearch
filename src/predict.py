# -*- coding: utf-8 -*-

import numpy as np
from keras.applications.vgg16 import preprocess_input, VGG16
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


def _similar(vec, knn, filenames):
    dist, indices = knn.kneighbors(vec.reshape(1, -1), n_neighbors=6)
    dist, indices = dist.flatten(), indices.flatten()
    return [filenames[indices[i]] for i in range(len(indices))]


def load_predictor(owner_id):
    filenames = open(config.images_order(owner_id), 'r').readline().split(',')
    vecs = load_sparse_matrix(config.vectors_path(owner_id))
    model = VGG16(include_top=False, weights='imagenet', pooling='max')
    knn = NearestNeighbors(metric='cosine', algorithm='brute')
    knn.fit(vecs)

    def similarity(file_path):
        vec = _vectorize(file_path, model)
        return _similar(vec, knn, filenames)

    return similarity
