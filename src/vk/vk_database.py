# -*- coding: utf-8 -*-

import pandas as pd

from vk.settings import config


def pictures_repo(owner_id):
    df = pd.read_csv(config.info_path(owner_id))

    def get_path(small_path):
        fname = small_path.split('/')[-1]
        parts = fname.split('_')
        id, date = int(parts[0]), int(parts[1])
        res = df[(df['id'] == id) & (df['date'] == date)]
        if res.shape[0] != 0:
            post = '{}_{}'.format(res['owner_id'].values[0], res['id'].values[0])
            return res['big'].values[0], vk_link(post)
        else:
            return

    return get_path


def vk_link(post):
    return "https://vk.com/feed?z=photo{}".format(post)
