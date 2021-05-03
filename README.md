# Camalian-Go

[![Go Report Card](https://goreportcard.com/badge/github.com/nazarhussain/camalian-go)](https://goreportcard.com/report/github.com/nazarhussain/camalian-go) [![GitHub license](https://img.shields.io/github/license/nazarhussain/camalian-go)](https://github.com/nazarhussain/camalian-go/blob/main/LICENSE)
 [![Build](https://github.com/nazarhussain/camalian-go/actions/workflows/build.yml/badge.svg)](https://github.com/nazarhussain/camalian-go/actions/workflows/build.yml) [![codecov](https://codecov.io/gh/nazarhussain/camalian-go/branch/main/graph/badge.svg?token=PZIHW9TIUJ)](https://codecov.io/gh/nazarhussain/camalian-go)

Color extraction utility in Golang

For details information on the utility please look at the following article. [Advance Color Quantization with Camalian](https://basicdrift.com/advance-color-quantization-with-camalian-ff369b65bb40). This implementation in Golang inspired by [Camalian Ruby Implementation](https://github.com/nazarhussain/camalian).


## Usage

### Binary Usage
To use the binary binary for utility. Install it with: 

```
$ GO111MODULE=on go get -u "github.com/nazarhussain/camalian-go/cli"
```

The usage is as follows:

```
Usage:
   camalian <path> {flags}

Commands: 
   help                          displays usage information
   version                       displays version number

Arguments: 
   path                          Path to the image file

Flags: 
   -h, --help                    displays usage information of the application or a command (default: false)
   -n, --number                  Number of colors to extract (default: 15)
   -q, --quantizer               Quantizer to use. histogram, kmeans, median_cut (default: histogram)
   -v, --version                 displays version number (default: false)
```

### Library Usage

Mention the package as your project mode dependencies.

```go
import (
  camalian "github.com/nazarhussain/camalian-go"
  quantizers "github.com/nazarhussain/camalian-go/quantizers"
```

Then use as following:

```go
path := "path/to/image/file"
palette := new(camalian.Palette).BuildFromImage(&camalian.Image{FilePath: path})
palette = palette.Quantize(quantizers.Histogram{}, 10)

fmt.Println(palette.ToHex())
```

Available Quantizers are `quantizers.Histogram{}`, `quantizers.KMeans{}` and `quantizers.MedianCut{}`.

For more details on each quantizers please [check this section](https://github.com/nazarhussain/camalian#quantization-algorithms).

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
