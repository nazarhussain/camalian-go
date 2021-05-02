package quantizers

import (
	"testing"

	"github.com/nazarhussain/camalian-go"
	"github.com/stretchr/testify/assert"
)

func TestKMeans(t *testing.T) {
	palleteColors := []string{
		"#4DD915", "#49CC23", "#45C031", "#41B43F", "#3DA84D", "#399C5B", "#359069", "#318478", "#2D7886", "#296C94", "#2560A2", "#2154B0", "#1D48BE", "#193CCC", "#1530DB"}

	q := KMeans{}
	filePath := "../testdata/palette.png"
	image := &camalian.Image{FilePath: filePath}
	palette := new(camalian.Palette).BuildFromImage(image)
	palette2 := q.Quantize(palette, 15)

	assert.Len(t, *palette2, 15)
	assert.Equal(t, palleteColors, palette2.ToHex())
}

func TestKMeansDistintColors(t *testing.T) {
	q := KMeans{}
	colors := []string{"#FF0000", "#00FF00", "#0000FF"}
	p1, _ := new(camalian.Palette).BuildFromHex(colors...)
	p2 := q.Quantize(p1, 3)

	assert.Len(t, *p2, 3)
	assert.Equal(t, colors, p2.ToHex())
}

func TestKMeansDistinctColorsLesserThanPixels(t *testing.T) {
	q := KMeans{}
	colors := []string{"#FF0000", "#00FF00", "#0000FF"}
	p1, _ := new(camalian.Palette).BuildFromHex(colors...)
	p2 := q.Quantize(p1, 2)

	assert.Len(t, *p2, 2)
}

func TestKMeansDistinctColorsMoreThanPixels(t *testing.T) {
	q := KMeans{}
	colors := []string{"#FF0000", "#00FF00", "#0000FF"}
	p1, _ := new(camalian.Palette).BuildFromHex(colors...)
	p2 := q.Quantize(p1, 4)

	assert.Len(t, *p2, 3)
}

func TestKMeansSameColors(t *testing.T) {
	q := KMeans{}
	colors := []string{"#FF0000", "#FF0000", "#FF0000"}
	p1, _ := new(camalian.Palette).BuildFromHex(colors...)
	p2 := q.Quantize(p1, 3)

	assert.Len(t, *p2, 1)
	assert.Equal(t, []string{"#FF0000"}, p2.ToHex())
}
