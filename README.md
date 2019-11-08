# Kawaii Search Service

## TODO

- exe to prepare dataset of photos
	- iterate over directory of images
	- convert image into embedding using NeuralNet
	- store all embeddings into a single binary file
- lib to similarity search
	- load emdebbings database into memory
	- convert image into embedding using NeuralNet
	- kNN to search similar images in the embeddings space
	- return top-K closest image IDs with distances
- exe to run API server
	- get image from the user's request
	- call similarity search
	- return results in json format
- exe to run telegram bot
	- get image from the user's message
	- call similarity search
	- send top-3 images with distance in the caption to the user
- exe to run web server
	- host html form to get image from the user in browser
	- send request to the API
	- show results in te galery
