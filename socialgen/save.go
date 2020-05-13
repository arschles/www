package main

import (
	"image"
	"image/png"
	"os"
)

func SavePNG(img image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
