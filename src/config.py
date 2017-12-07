# -*- coding: utf-8 -*-

pool_size = 20

def images_path(onwer_id):
    return 'data/images_{}'.format(onwer_id)


def images_glob_path(onwer_id):
    return 'data/images_{}/*.jpg'.format(onwer_id)


def images_order(onwer_id):
    return 'submission/images_order_{}.csv'.format(onwer_id)


def info_path(onwer_id):
    return 'data/photos_{}.csv'.format(onwer_id)


def vectors_path(onwer_id):
    return 'submission/images_vec_{}.npz'.format(onwer_id)
