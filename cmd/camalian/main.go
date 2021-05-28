package main

import (
	"fmt"
	"strings"

	"github.com/nazarhussain/camalian-go"
	quantizers "github.com/nazarhussain/camalian-go/pkg/quantizers"
	"github.com/thatisuday/commando"
)

var availableQuantizers = map[string]camalian.Quantizer{
	"histogram":  quantizers.Histogram{},
	"kmeans":     quantizers.KMeans{},
	"median_cut": quantizers.MedianCut{},
}

func main() {
	commando.
		SetExecutableName("camalian").
		SetVersion("v0.2.0").
		SetDescription("Extract color palettes from images.")

	commando.
		Register(nil).
		AddArgument("path", "Path to the image file", "").
		AddFlag("quantizer,q", "Quantizer to use. histogram, kmeans, median_cut", commando.String, "histogram").
		AddFlag("number,n", "Number of colors to extract", commando.Int, 15).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			q := flags["quantizer"].Value.(string)
			n := flags["number"].Value.(int)

			if q != "histogram" && q != "kmeans" && q != "median_cut" {
				fmt.Println("Quantizer value can be one of \"histogram, kmeans, median_cut\"")
				return
			}

			if n < 1 {
				fmt.Println("Please specify number of colors larger than zero")
				return
			}

			path := args["path"].Value
			palette := new(camalian.Palette).BuildFromImage(&camalian.Image{FilePath: path})
			palette = palette.Quantize(availableQuantizers[q], uint(n))

			fmt.Println(strings.Join(palette.ToHex(), ","))
		})

	// parse command-line arguments from the STDIN
	commando.Parse(nil)
}
