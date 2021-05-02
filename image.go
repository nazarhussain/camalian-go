package camalian

import (
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Image struct {
	FilePath string
}

func (i *Image) Pixels() []RGB {
	existingImageFile, err := os.Open(i.FilePath)

	if err != nil {
		panic(err)
	}

	defer existingImageFile.Close()

	imageData, _, err := image.Decode(existingImageFile)
	if err != nil {
		panic(err)
	}

	bounds := imageData.Bounds()

	var colors = make([]RGB, (bounds.Max.X * bounds.Max.Y))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := imageData.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 8 reduces this to the range [0, 255].
			colors[x+(y*bounds.Max.X)] = RGB{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8)}
		}
	}

	return colors
}
