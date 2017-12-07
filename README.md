# Images similarity

This is a demo of applying VGG16 and kNN to build similar image search. 

## Dataset

__[download my dataset](https://yadi.sk/d/lNREFvKa3QPSRs)__

You can use any big dataset with images. I used cosplay pictures from this group: [vk/wtfcute](https://vk.com/wtfcute).

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

## TODO

- [ ] create a single main file to do all the steps
- [ ] write a blog post about this
- [ ] create an app to find similar photos from the dataset by an image URL
- [ ] create web service for telegram chat
- [ ] download more images 
