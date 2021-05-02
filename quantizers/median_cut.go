package quantizers

import (
	"reflect"

	"github.com/nazarhussain/camalian-go"
	utils "github.com/nazarhussain/camalian-go/utils"
)

type MedianCut struct{}

type splitInfo struct {
	colorRange uint8
	groupIndex int
	color      string
}

func (MedianCut) Quantize(p *camalian.Palette, count uint) *camalian.Palette {
	p2 := p.Unique()

	if len(*p) <= int(count) {
		return p2.SortSimilarColors()
	}

	groups := []*camalian.Palette{p2}
	limit, _ := utils.FindMinMax(int(count), len(*p))

	for {
		split := determineGroupSplit(groups)
		if split.groupIndex != -1 {
			group1, group2 := splitGroup(groups[split.groupIndex], split)
			groups = append(groups[:split.groupIndex], groups[split.groupIndex+1:]...)
			groups = append(groups, group1)
			groups = append(groups, group2)
		}

		if len(groups) >= limit.(int) {
			break
		}
	}

	result := make(camalian.Palette, 0, count)

	for i, v := range groups {
		if uint(i) == count {
			break
		}

		result = append(result, v.AverageColor())
	}
	result.SortSimilarColors()

	return &result
}

func determineGroupSplit(groups []*camalian.Palette) splitInfo {
	split := splitInfo{groupIndex: -1}

	for index, group := range groups {
		if len(*group) == 0 {
			continue
		}

		colors := make([]interface{}, len(*group))
		for index, value := range *group {
			colors[index] = value
		}

		reds := utils.Collect(colors, func(v interface{}) interface{} {
			return v.(*camalian.Color).R
		})
		minRed, maxRed := utils.FindMinMax(reds...)
		redRange := maxRed.(uint8) - minRed.(uint8)

		greens := utils.Collect(colors, func(v interface{}) interface{} {
			return v.(*camalian.Color).G
		})
		minGreen, maxGreen := utils.FindMinMax(greens...)
		greenRange := maxGreen.(uint8) - minGreen.(uint8)

		blues := utils.Collect(colors, func(v interface{}) interface{} {
			return v.(*camalian.Color).B
		})
		minBlue, maxBlue := utils.FindMinMax(blues...)
		blueRange := maxBlue.(uint8) - minBlue.(uint8)

		newSplit := splitInfo{groupIndex: -1}
		if redRange > greenRange && redRange > blueRange {
			newSplit = splitInfo{groupIndex: index, colorRange: (maxRed.(uint8) - minRed.(uint8)), color: "R"}
		} else if greenRange > redRange && greenRange > blueRange {
			newSplit = splitInfo{groupIndex: index, colorRange: (maxGreen.(uint8) - minGreen.(uint8)), color: "G"}
		} else {
			newSplit = splitInfo{groupIndex: index, colorRange: (maxBlue.(uint8) - minBlue.(uint8)), color: "B"}
		}

		if split.groupIndex == -1 {
			split = newSplit
			continue
		}

		if newSplit.groupIndex != -1 && newSplit.colorRange > split.colorRange {
			split = newSplit
		}
	}

	return split
}

func getField(v *camalian.Color, field string) uint8 {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return uint8(f.Uint())
}

func splitGroup(group *camalian.Palette, split splitInfo) (*camalian.Palette, *camalian.Palette) {
	coloSortKey := map[string]camalian.SortProperty{"R": camalian.SortProperty(camalian.Red), "G": camalian.SortProperty(camalian.Green), "B": camalian.SortProperty(camalian.Blue)}
	colors := group.SortBy(coloSortKey[split.color], true)

	if len(*colors) == 2 {
		p1, _ := new(camalian.Palette).BuildFromHex((*colors)[0].ToHex())
		p2, _ := new(camalian.Palette).BuildFromHex((*colors)[1].ToHex())

		return p1, p2
	}

	medianIndex := len(*colors) / 2
	medianValue := getField((*colors)[medianIndex], split.color)

	group1 := new(camalian.Palette)
	group2 := new(camalian.Palette)

	for _, color := range *colors {
		if getField(color, split.color) < medianValue {
			*group1 = append(*group1, color)
		} else {
			*group2 = append(*group2, color)
		}
	}

	return group1, group2
}
