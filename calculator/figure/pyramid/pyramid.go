package pyramid

type Pyramid struct {
	baseArea float64
	height   float64
}

func NewPyramid(baseArea, height float64) *Pyramid {
	if baseArea <= 0 || height <= 0 {
		return nil
	}

	return &Pyramid{
		baseArea: baseArea,
		height:   height,
	}
}

func (p *Pyramid) Volume() float64 {
	if p == nil {
		return 0
	}

	return 1.0 / 3.0 * p.baseArea * p.height
}
