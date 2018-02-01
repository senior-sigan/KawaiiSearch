# Kawaii Search (Images similarity)

The blog post describing how it works is [here](https://medium.com/@senior_sigan/similar-images-search-ce433491059b).

This is a demo of applying VGG16 and kNN to build similar image search. 

![vk.com/tokyofashion](https://i.imgur.com/e9bpwWY.png)
![vk.com/tokyofashion](https://i.imgur.com/dDAJCuY.png)

## Dataset

You can use any big dataset with images. I used pictures from this group about fashion: [vk/tokyofashion](https://vk.com/tokyofashion).

```
./data
-- photos.csv # is a csv file with pictures' info and url.
-- images     # is a directory with pictures
```

You can use `src/get_images.py` to get all pictures info and urls from a specified group in the vk. Use `config.py` to set vk group id.

You can use `src/download_images.py` to download images listed in the `data/photos.csv`.

## Training

I use pretrained VGG16 from the keras, but without last layers, only global max pooling. So i get 512 features per image, that i used in the kNN with the cosine metric to calculate similarity.

But you have to generate all the 512-sized vectors for each image, so run `src/vectorize_image.py` to do it. On the GPU Tesla K80 in the gcloud for 50_000 images it takes 20 minutes. The result will be saved in the `submission/images_vec.npz` with `submissions/images_order.csv`.

## Evaluating

Look into `test.ipynb` file for example of using this model.

## Deploy

Modify and copy app.service to the `/etc/systemd/system`. Run `systemctl daemon-reload` and `systemctl enable app.service`.

## TODO

- [x] create a single main file to do all the steps
- [x] write a blog post about this
- [x] build a web service with
- [ ] add feature to find similar photos by an image URL
- [ ] create web service for telegram chat
- [x] download more images 
