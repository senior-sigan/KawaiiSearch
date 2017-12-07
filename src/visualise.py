# -*- coding: utf-8 -*-

from PIL import Image
from matplotlib import pyplot as plt


def draw(fnames, origin=None):
    if origin is not None:
        plt.imshow(Image.open(origin))
        plt.axis('off')
    plt.figure(figsize=(30, 15))
    for i in range(len(fnames)):
        f, d = fnames[i]
        try:
            img = Image.open(f)
            plt.subplot(1, 10, i + 1)
            plt.axis('off')
            plt.title("{0:.4f}".format(d))
            plt.imshow(img)
        except Exception as e:
            print(e)
