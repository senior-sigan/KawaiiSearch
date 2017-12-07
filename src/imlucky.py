# -*- coding: utf-8 -*-

from database import pictures_repo
from predict_dummy import load_predictor, random

owner_id = -70232735

print("Loading imlucky")
pred = load_predictor(owner_id)
rf = random(owner_id)
repo = pictures_repo(owner_id)
print("Imlucky loaded")

def imlucky():
    f = rf()
    files = pred(f, 12)
    return [(repo(file), d) for file, d in files]
