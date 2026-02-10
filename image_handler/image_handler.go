package image_handler

import (
	"image"

	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/imgio"
)

func SavePNG(filename string, blurred *image.RGBA) bool {
	if err := imgio.Save(filename, blurred, imgio.PNGEncoder()); err != nil {
		return false
	}
	return true
}

func SaveJPEG(filename string, blurred *image.RGBA) bool {
	if err := imgio.Save(filename, blurred, imgio.JPEGEncoder(100)); err != nil {
		return false
	}
	return true
}

func GaussianBlur(image image.Image, blur_radius float64) *image.RGBA {
	return blur.Gaussian(image, blur_radius)
}
