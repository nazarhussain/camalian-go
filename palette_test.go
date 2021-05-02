package camalian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaletteBuildFromHex(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#FF0000", true)
	c2, _ := new(Color).BuildFromHex("#00FF00", true)

	p, err := new(Palette).BuildFromHex("#FF0000", "#00FF00")

	assert.NoError(t, err)
	assert.Len(t, *p, 2)
	assert.Equal(t, (*p)[0], c1)
	assert.Equal(t, (*p)[1], c2)
}

func TestPaletteBuildFromImage(t *testing.T) {
	filePath := "testdata/palette.png"
	var image *Image = &Image{FilePath: filePath}

	p := new(Palette).BuildFromImage(image)

	assert.Len(t, *p, 54000)
}

func TestPaletteSortByRGBAsc(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#00FF00", true)
	p, _ := new(Palette).BuildFromHex("#00FF00", "#FF0000")
	p.SortBy(SortProperty(Green), true)

	assert.Equal(t, (*p)[1], c1)
}

func TestPaletteSortByRGBDesc(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#00FF00", true)
	p, _ := new(Palette).BuildFromHex("#FF0000", "#00FF00")
	p.SortBy(SortProperty(Green), false)

	assert.Equal(t, (*p)[0], c1)
}

func TestPaletteSortByHSLAsc(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#00FF00", true)
	p, _ := new(Palette).BuildFromHex("#00FF00", "#FF0000")
	p.SortBy(SortProperty(HSLHue), true)

	assert.Equal(t, (*p)[1], c1)
}

func TestPaletteSortByHSLDesc(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#00FF00", true)
	p, _ := new(Palette).BuildFromHex("#00FF00", "#FF0000")
	p.SortBy(SortProperty(HSLHue), false)

	assert.Equal(t, (*p)[0], c1)
}

func TestPaletteSortSimilarColors(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#00FF00", true)
	c2, _ := new(Color).BuildFromHex("#009900", true)
	c3, _ := new(Color).BuildFromHex("#FF0000", true)
	c4, _ := new(Color).BuildFromHex("#550000", true)
	c5, _ := new(Color).BuildFromHex("#0000FF", true)
	c6, _ := new(Color).BuildFromHex("#000055", true)

	p := Palette{c1, c3, c5, c2, c4, c6}
	p.SortSimilarColors()

	assert.Equal(t, []string{"#FF0000", "#550000", "#00FF00", "#009900", "#0000FF", "#000055"}, p.ToHex())
}

func TestAverageColor(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#FF0000", true)
	c2, _ := new(Color).BuildFromHex("#FF00FF", true)

	p := Palette{c1, c2}
	avg := p.AverageColor()

	expected, _ := new(Color).BuildFromHex("#FF007f", true)

	assert.Equal(t, expected, avg)
}

func TestLightColors(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#000000", true)
	c2, _ := new(Color).BuildFromHex("#666666", true)
	c3, _ := new(Color).BuildFromHex("#999999", true)

	p := Palette{c1, c2, c3}
	p2 := p.LightColors(30, 50)

	assert.Equal(t, []string{"#666666"}, p2.ToHex())
}

func TestUniqueColors(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#FF0000", true)
	c2, _ := new(Color).BuildFromHex("#FF0000", true)
	c3, _ := new(Color).BuildFromHex("#999999", true)

	p := Palette{c1, c2, c3}
	p2 := p.Unique()

	assert.Equal(t, []string{"#FF0000", "#999999"}, p2.ToHex())
}

func TestCommonColors(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#FF0000", true)
	c2, _ := new(Color).BuildFromHex("#00FF00", true)
	c3, _ := new(Color).BuildFromHex("#999999", true)

	p := Palette{c1, c2, c3}
	p2 := Palette{c1, c2}
	p3 := p.Common(&p2)

	assert.Equal(t, []string{"#FF0000", "#00FF00"}, p3.ToHex())
}
