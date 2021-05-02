package camalian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImagePixelsRead(t *testing.T) {
	filePath := "testdata/palette.png"

	var image *Image = &Image{FilePath: filePath}
	pixels := image.Pixels()

	assert.Len(t, pixels, (900 * 60))
	assert.Equal(t, pixels[0], RGB{R: 21, G: 48, B: 219})
	assert.Equal(t, pixels[len(pixels)-1], RGB{R: 77, G: 217, B: 21})
}
