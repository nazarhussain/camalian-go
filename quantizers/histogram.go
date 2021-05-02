package quantizers

import (
	"fmt"
	"math"
	"sort"

	"github.com/nazarhussain/camalian-go"
)

type Histogram struct{}

const maxRetry = 5

func (Histogram) Quantize(p *camalian.Palette, count uint) *camalian.Palette {
	var retryCount uint = 0
	var p2 *camalian.Palette

	for {
		p2 = process(p, (count + retryCount))

		if uint(len(*p2)) >= count || retryCount == maxRetry {
			break
		}

		retryCount += 1
	}

	return p2
}

func process(p *camalian.Palette, count uint) *camalian.Palette {
	bucketSize := uint(math.Ceil(255 / float64(count)))
	buckets := make(map[string]camalian.Palette)

	for _, color := range *p {
		key := bucketKey(color, bucketSize)

		if buckets[key] != nil {
			buckets[key] = make(camalian.Palette, 0)
		}
		buckets[key] = append(buckets[key], color)
	}

	allBuckets := make([]camalian.Palette, 0, len(buckets))
	for _, v := range buckets {
		allBuckets = append(allBuckets, v)
	}

	sort.SliceStable(allBuckets, func(i, j int) bool {
		return len(allBuckets[i]) < len(allBuckets[j])
	})

	result := make(camalian.Palette, 0, count)

	for i, v := range allBuckets {
		if uint(i) == count {
			break
		}

		result = append(result, v.AverageColor())
	}
	result.SortSimilarColors()

	return &result

}

func bucketKey(c *camalian.Color, bucketSize uint) string {
	size := uint8(bucketSize)

	return fmt.Sprintf("%s:%s:%s", colorKey(c.R, size), colorKey(c.G, size), colorKey(c.B, size))
}

func colorKey(color, bucketSize uint8) string {
	if color == 0 {
		return "0"
	}

	return fmt.Sprintf("%d", (color-1)/bucketSize)
}
