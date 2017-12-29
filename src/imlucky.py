# -*- coding: utf-8 -*-

from predict_dummy import load_predictor, random


def load(dir_name, pictures_repo):
    print("Loading imlucky")
    pred = load_predictor(dir_name)
    rf = random(dir_name)
    repo = pictures_repo(dir_name)
    print("Imlucky loaded")

    def imlucky():
        f = rf()
        files = pred(f, 12)
        return [(repo(file), d) for file, d in files]

    return imlucky
