# -*- coding: utf-8 -*-

pool_size = 20


def images_path(dir_name):
    return 'data/{}'.format(dir_name)


def images_glob_path(dir_name):
    return 'data/{}/*.jpg'.format(dir_name)


def images_order(dir_name):
    return 'submission/images_order_{}.csv'.format(dir_name)

def vectors_path(dir_name):
    return 'submission/images_vec_{}.npz'.format(dir_name)
