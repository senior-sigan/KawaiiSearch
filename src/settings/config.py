# -*- coding: utf-8 -*-

pool_size = 20


def images_path(owner_id):
    return 'data/images_{}'.format(owner_id)


def images_glob_path(owner_id):
    return 'data/images_{}/*.jpg'.format(owner_id)


def images_order(owner_id):
    return 'submission/images_order_{}.csv'.format(owner_id)


def info_path(owner_id):
    return 'data/photos_{}.csv'.format(owner_id)


def vectors_path(owner_id):
    return 'submission/images_vec_{}.npz'.format(owner_id)
