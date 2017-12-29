# -*- coding: utf-8 -*-

import numpy as np
from sklearn.neighbors import NearestNeighbors

from settings import config
from sparce import load_sparse_matrix


def _similar(vec, knn, filenames, n_neighbors=6):
    dist, indices = knn.kneighbors(vec.reshape(1, -1), n_neighbors=n_neighbors)
    dist, indices = dist.flatten(), indices.flatten()
    return [(filenames[indices[i]], dist[i]) for i in range(len(indices))]


def load_predictor(dir_name):
    print("load predictor")
    filenames = open(config.images_order(dir_name), 'r').readline().split(',')
    vecs = load_sparse_matrix(config.vectors_path(dir_name))
    knn = NearestNeighbors(metric='cosine', algorithm='brute')
    knn.fit(vecs)

    def similarity(vec, n_neighbors=6):
        return _similar(vec, knn, filenames, n_neighbors)

    print("Predictor loaded")
    return similarity


def random(dir_name):
    print("Preparing random generator")
    vecs = load_sparse_matrix(config.vectors_path(dir_name)).toarray()
    s = vecs.shape[0]

    def rf():
        return vecs[np.random.randint(0, s)]

    print("Random generator is ready")
    return rf
