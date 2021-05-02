package quantizers

import (
	"math/rand"
	"time"

	"github.com/nazarhussain/camalian-go"
	utils "github.com/nazarhussain/camalian-go/utils"
)

type KMeans struct{}

func (KMeans) Quantize(p *camalian.Palette, count uint) *camalian.Palette {
	p2 := p.Unique()

	if len(*p2) <= int(count) {
		return p2.SortSimilarColors()
	}

	means := initialMeans(p2, count)

	for {
		newMeans := distanceWithAverage(p2, means)
		common := means.Common(newMeans)

		if len(*common) == len(*means) {
			break
		}

		means = newMeans
	}

	return means.SortSimilarColors()
}

func initialMeans(p *camalian.Palette, count uint) *camalian.Palette {
	c := int(count)

	if c > len(*p) {
		c = len(*p)
	}

	means := new(camalian.Palette)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(*means) != c {
		for i := 0; i < (c - len(*means)); i += 1 {
			*means = append(*means, (*p)[r.Intn(len(*p))])
		}
		means = means.Unique()
	}
	return means
}

func distanceWithAverage(p, means *camalian.Palette) *camalian.Palette {
	groups := make(map[int]*camalian.Palette)

	for _, color := range *p {
		distances := make([]interface{}, 0, len(*means))

		for _, m := range *means {
			distances = append(distances, m.RgbDistance(color))
		}

		min, _ := utils.FindMinMax(distances...)
		minDistanceIndex := utils.Find(distances, func(value interface{}) bool {
			return value.(float64) == min
		})

		if _, value := groups[minDistanceIndex]; !value {
			groups[minDistanceIndex] = new(camalian.Palette)
			*groups[minDistanceIndex] = append(*groups[minDistanceIndex], color)
		}
	}

	result := new(camalian.Palette)
	for _, val := range groups {
		*result = append(*result, val.AverageColor())
	}

	return result
}
