# -*- coding: utf-8 -*-

import os
import sys
from multiprocessing.dummy import Pool

import requests

import config

KEYS = ['id', 'owner_id', 'album_id', 'date', 'big', 'small']


def main(owner_id):
    root = config.images_path(owner_id)
    os.makedirs(root, exist_ok=True)
    count = 0

    def _download(url):
        nonlocal count
        count += 1
        if count % 50 == 0:
            print("{}".format(count))
        try:
            download(root, url)
        except Exception as e:
            print(e)

    Pool(config.pool_size).map(_download, data(config.info_path(owner_id)))


def download(root, d):
    url = d['small']
    fname = url.split('/')[-1]
    name = "{}_{}_{}".format(d['id'], d['date'], fname)
    with open(os.path.join(root, name), 'wb') as file:
        res = requests.get(url)
        file.write(res.content)


def data(path):
    with open(path, "r") as fd:
        fd.readline()
        while True:
            line = fd.readline()
            if line is None or len(line) == 0:
                break
            line = line.split('\n')[0].split(',')
            d = {}
            for i in range(len(KEYS)):
                d[KEYS[i]] = line[i]
            yield d


if __name__ == '__main__':
    if (len(sys.argv) != 2):
        print("Should be `python3 src/download_images.py GROUP_ID`")
    main(sys.argv[1])
