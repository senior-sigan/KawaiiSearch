package images

import (
	"bytes"
	"errors"
	"image"
	"io/ioutil"

)

func DecodeImage(raw []byte) (image.Image, error) {
	if len(raw) == 0 {
		return nil, errors.New("image is empty")
	}
	img, _, err := Decode(bytes.NewReader(raw))
	return img, err
}

func ReadImage(filename string) (image.Image, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return DecodeImage(file)
}
