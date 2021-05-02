package camalian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	c := new(Color).Build(10, 20, 30, false)

	assert.Equal(t, c.R, uint8(10))
	assert.Equal(t, c.G, uint8(20))
	assert.Equal(t, c.B, uint8(30))
	assert.Equal(t, c.HSL, HSL{})
	assert.Equal(t, c.HSV, HSV{})
}

func TestBuildWihComponets(t *testing.T) {
	c := new(Color).Build(10, 20, 30, true)

	assert.Equal(t, c.R, uint8(10))
	assert.Equal(t, c.G, uint8(20))
	assert.Equal(t, c.B, uint8(30))
	assert.Equal(t, c.HSL, HSL{H: 210, S: 50, L: 7.8431377})
	assert.Equal(t, c.HSV, HSV{H: 210, S: 66.666664, V: 11.764706})
}

func TestBuildFromHex(t *testing.T) {
	c, _ := new(Color).BuildFromHex("#0a141e", false)

	assert.Equal(t, c.R, uint8(10))
	assert.Equal(t, c.G, uint8(20))
	assert.Equal(t, c.B, uint8(30))
	assert.Equal(t, c.HSL, HSL{})
	assert.Equal(t, c.HSV, HSV{})
}

func TestBuildFromHexWithComponents(t *testing.T) {
	c, _ := new(Color).BuildFromHex("#0a141e", true)

	assert.Equal(t, c.R, uint8(10))
	assert.Equal(t, c.G, uint8(20))
	assert.Equal(t, c.B, uint8(30))
	assert.Equal(t, c.HSL, HSL{H: 210, S: 50, L: 7.8431377})
	assert.Equal(t, c.HSV, HSV{H: 210, S: 66.666664, V: 11.764706})
}

func TestBuildFromHexWithoutHash(t *testing.T) {
	c, _ := new(Color).BuildFromHex("0a141e", false)

	assert.Equal(t, c.R, uint8(10))
	assert.Equal(t, c.G, uint8(20))
	assert.Equal(t, c.B, uint8(30))
}

func TestBuildFromHexWithInvalidHash(t *testing.T) {
	_, e := new(Color).BuildFromHex("#0ag41e", false)

	assert.Error(t, e)
}

func TestRgbDistance(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#0a241e", false)
	c2, _ := new(Color).BuildFromHex("#0a142e", false)
	rgbDistance := c1.RgbDistance(c2)

	assert.Equal(t, 22.627416997969522, rgbDistance)
}

func TestHuDistance(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#0a241e", true)
	c2, _ := new(Color).BuildFromHex("#0a142e", true)
	hueDistance := c1.HueDistance(c2)

	assert.Equal(t, hueDistance, uint16(57))
}

func TestToHex(t *testing.T) {
	c1, _ := new(Color).BuildFromHex("#ff0005", false)

	assert.Equal(t, c1.ToHex(), "#FF0005")
}
