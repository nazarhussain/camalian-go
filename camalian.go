package camalian

// The interface to fetch any quantizeable object
type Quantizeable interface {
	Pixels() []RGB
}

// The interface to to quantizer algorithm
type Quantizer interface {
	Quantize(*Palette, uint) *Palette
}
