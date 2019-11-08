package main

import (
	"image"
	_ "image/jpeg"

	"github.com/disintegration/imaging"
	"github.com/sirupsen/logrus"
)

type RGB struct {
	R, G, B float32
}

func flattenImg(img image.Image, mean RGB, std RGB) []float32 {
	width := 224
	height := 224
	size := width * height
	res := make([]float32, 3*width*height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := img.At(x, y)
			r, g, b, _ := p.RGBA()

			rn := float32(r) / 0xffff
			gn := float32(g) / 0xffff
			bn := float32(b) / 0xffff

			res[x+y*width] = (rn - mean.R) / std.R
			res[x+y*width+size] = (gn - mean.G) / std.G
			res[x+y*width+size*2] = (bn - mean.B) / std.B
		}
	}
	return res
}

// ProcessImage converts image into single-dimensional vector in format RRR...GGG...BBB
// Each color chanel is normalized using ImageNet mean and std
// as described https://pytorch.org/docs/stable/torchvision/models.html
func ProcessImage(img image.Image, width int, height int) []float32 {
	logrus.Infof("Image size %v", img.Bounds())
	img224 := imaging.Resize(img, width, height, imaging.Lanczos)
	logrus.Infof("Image size %v", img224.Bounds())
	rawImg := flattenImg(img224, RGB{0.485, 0.456, 0.406}, RGB{0.229, 0.224, 0.225})
	return rawImg
}
