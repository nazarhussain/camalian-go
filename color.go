package camalian

import (
	"fmt"
	"math"
	"strconv"

	utils "github.com/nazarhussain/camalian-go/pkg/utils"
)

type HSL struct {
	H uint16
	S float32
	L float32
}

type HSV struct {
	H uint16
	S float32
	V float32
}

type RGB struct {
	R uint8
	G uint8
	B uint8
}

type Color struct {
	RGB
	HSV HSV
	HSL HSL
}

// Build Color object from rgb values
// Specify buildAll = true if you want to calculate HSL/HSV values as well
func (c *Color) Build(r, g, b uint8, buildAll bool) *Color {
	c.R = r
	c.G = g
	c.B = b

	if buildAll {
		c.buildComponents()
	}
	return c
}

// Build Color object from Hex string e.g. #FF7913
// Specify buildAll = true if you want to calculate HSL/HSV values as well
func (c *Color) BuildFromHex(hex string, buildAll bool) (*Color, error) {
	index := 0

	if hex[0] == '#' {
		index = 1
	}

	r, err := strconv.ParseUint(hex[index:index+2], 16, 16)
	if err != nil {
		return nil, err
	}
	c.R = uint8(r)
	index += 2

	g, err := strconv.ParseUint(hex[index:index+2], 16, 16)
	if err != nil {
		return nil, err
	}
	c.G = uint8(g)
	index += 2

	b, err := strconv.ParseUint(hex[index:index+2], 16, 16)
	if err != nil {
		return nil, err
	}
	c.B = uint8(b)

	if buildAll {
		c.buildComponents()
	}
	return c, nil
}

// Calculate Hue distance between two colors
func (c *Color) HueDistance(c2 *Color) uint16 {
	h1 := ((c.HSL.H - c2.HSL.H) % 360)
	h2 := ((c2.HSL.H - c.HSL.H) % 360)
	min, _ := utils.FindMinMax(h1, h2)

	return min.(uint16)
}

// Calculate RGB distance between two colors
func (c *Color) RgbDistance(c2 *Color) float64 {
	rsqr := math.Pow((float64(c.R) - float64(c2.R)), 2)
	gsqr := math.Pow((float64(c.G) - float64(c2.G)), 2)
	bsqr := math.Pow((float64(c.B) - float64(c2.B)), 2)

	return math.Sqrt(rsqr + gsqr + bsqr)
}

func (c *Color) ToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func (c *Color) buildComponents() {
	ri := float32(c.R) / 255.0
	gi := float32(c.G) / 255.0
	bi := float32(c.B) / 255.0
	hsl := HSL{}
	hsv := HSV{}

	min, max := utils.FindMinMax(ri, gi, bi)
	cmin := min.(float32)
	cmax := max.(float32)
	delta := cmax - cmin

	hsl.L = float32(cmax+cmin) / 2.0

	if delta == 0 {
		hsl.H = 0
	} else if cmax == ri {
		hsl.H = 60 * (uint16(((gi - bi) / delta)) % 6)
	} else if cmax == gi {
		hsl.H = uint16(60 * (((bi - ri) / delta) + 2))
	} else if cmax == bi {
		hsl.H = uint16(60 * (((ri - gi) / delta) + 4))
	}

	if delta == 0 {
		hsl.S = 0
	} else {
		hsl.S = float32(delta) / float32(1-math.Abs(float64(2*hsl.L-1)))
	}

	hsl.S = hsl.S * 100
	hsl.L = hsl.L * 100

	// HSV Calculation
	// Hue calculation
	hsv.H = hsl.H
	if delta == 0 {
		hsv.H = 0
	} else if cmax == ri {
		hsv.H = 60 * (uint16((gi-bi)/delta) % 6)
	} else if cmax == gi {
		hsv.H = uint16(60 * (((bi - ri) / delta) + 2))
	} else if cmax == bi {
		hsv.H = uint16(60 * (((ri - gi) / delta) + 4))
	}

	// Saturation calculation
	if cmax == 0 {
		hsv.S = 0
	} else {
		hsv.S = float32(delta) / float32(cmax) * 100
	}

	// Value calculation
	hsv.V = float32(cmax) * 100

	c.HSL = hsl
	c.HSV = hsv
}
