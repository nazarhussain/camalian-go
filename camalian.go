package camalian

type Quantizeable interface {
	Pixels() []RGB
}

type Quantizer interface {
	Quantize(*Palette, uint) *Palette
}
