package camalian

import (
	"sort"

	utils "github.com/nazarhussain/camalian-go/utils"
)

type Palette []*Color

type SortProperty int

const (
	HSLHue SortProperty = iota
	HSLSaturation
	HSLLightness
	HSVHue
	HSVSaturation
	HSVValue
	Red
	Green
	Blue
)

var colorSortPropMap = map[SortProperty]func(*Color, *Color, bool) bool{
	HSLHue: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.HSL.H < c2.HSL.H
		}
		return c1.HSL.H > c2.HSL.H
	},
	HSLSaturation: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.HSL.S < c2.HSL.S
		}
		return c1.HSL.S > c2.HSL.S
	},
	HSLLightness: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.HSL.L < c2.HSL.L
		}
		return c1.HSL.L > c2.HSL.L
	},
	HSVHue: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.HSV.H < c2.HSV.H
		}
		return c1.HSV.H > c2.HSV.H
	},
	HSVSaturation: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.HSV.S < c2.HSV.S
		}
		return c1.HSV.S > c2.HSV.S
	},
	HSVValue: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.HSV.V < c2.HSV.V
		}
		return c1.HSV.V > c2.HSV.V
	},
	Red: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.R < c2.R
		}
		return c1.R > c2.R
	},
	Green: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.G < c2.G
		}
		return c1.G > c2.G
	},
	Blue: func(c1 *Color, c2 *Color, order bool) bool {
		if order {
			return c1.B < c2.B
		}
		return c1.B > c2.B
	},
}

// Build color palette from hex strings
//  new(Palette).BuildFromHex("#FF0000", "#00FF00")
func (p *Palette) BuildFromHex(colors ...string) (*Palette, error) {
	for _, hex := range colors {
		c, err := new(Color).BuildFromHex(hex, true)

		if err != nil {
			return nil, err
		}

		*p = append(*p, c)
	}

	return p, nil
}

// Build color palette from a Quantizeable interface
func (p *Palette) BuildFromImage(source Quantizeable) *Palette {
	for _, c := range source.Pixels() {
		c := new(Color).Build(c.R, c.G, c.B, true)

		*p = append(*p, c)
	}

	return p
}

// Sort color palette by some properties
func (p *Palette) SortBy(v SortProperty, order bool) *Palette {
	sort.SliceStable(*p, func(i, j int) bool {
		return colorSortPropMap[v]((*p)[i], (*p)[j], order)
	})

	return p
}

// Sort palette to visually similar colors
func (p *Palette) SortSimilarColors() *Palette {
	return p.SortBy(SortProperty(HSLHue), true)
}

// Calculate avargage color in RGB Color Space for the whole palette
func (p *Palette) AverageColor() *Color {
	colors := make([]interface{}, len(*p))
	for index, value := range *p {
		colors[index] = value
	}

	reds := utils.Sum(utils.Collect(colors, func(v interface{}) interface{} {
		return v.(*Color).R
	}))

	greens := utils.Sum(utils.Collect(colors, func(v interface{}) interface{} {
		return v.(*Color).G
	}))

	blues := utils.Sum(utils.Collect(colors, func(v interface{}) interface{} {
		return v.(*Color).B
	}))

	size := float64(len(*p))

	return new(Color).Build(uint8(reds.(float64)/size), uint8(greens.(float64)/size), uint8(blues.(float64)/size), true)
}

// Filter colors in particular range for lightness in HSL color space
func (p *Palette) LightColors(limit1, limit2 uint16) *Palette {
	newPalette := new(Palette)
	min, max := utils.FindMinMax(limit1, limit2)
	min1 := float32(min.(uint16))
	max1 := float32(max.(uint16))

	for _, c := range *p {
		if c.HSL.L > min1 && c.HSL.L < max1 {
			*newPalette = append(*newPalette, c)
		}
	}

	return newPalette
}

func (p *Palette) ToHex() []string {
	colors := []string{}

	for _, c := range *p {
		colors = append(colors, c.ToHex())
	}

	return colors
}

// Quantize the palette by specific Quantizer for given number of colors
func (p *Palette) Quantize(q Quantizer, count uint) *Palette {
	return q.Quantize(p, count)
}

// Remove redundant colors in palette and get unique colors
func (p *Palette) Unique() *Palette {
	keys := make(map[string]bool)
	p2 := new(Palette)

	for _, color := range *p {
		key := color.ToHex()

		if _, value := keys[key]; !value {
			keys[key] = true
			*p2 = append(*p2, color)
		}
	}

	return p2
}

// Fetch common colors among two color palettes
func (p *Palette) Common(p2 *Palette) *Palette {
	m := make(map[string]bool)
	result := new(Palette)

	for _, color := range *p {
		m[color.ToHex()] = false
	}

	for _, color := range *p2 {
		if _, v := m[color.ToHex()]; v {
			m[color.ToHex()] = true
		}
	}

	for hex, exists := range m {
		if exists {
			c, _ := new(Color).BuildFromHex(hex, true)
			*result = append(*result, c)
		}
	}

	return result.SortSimilarColors()
}
