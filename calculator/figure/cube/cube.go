package cube

import "math"

type Cube struct {
	side float64
}

func NewCube(side float64) *Cube {
	if side <= 0 {
		return nil
	}

	return &Cube{
		side: side,
	}
}

func (c *Cube) Volume() float64 {
	if c == nil {
		return 0
	}

	return math.Pow(c.side, 3)
}
