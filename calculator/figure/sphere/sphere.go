package sphere

import "math"

type Sphere struct {
	radius float64
}

func NewSphere(radius float64) *Sphere {
	if radius <= 0 {
		return nil
	}

	return &Sphere{
		radius: radius,
	}
}

func (s *Sphere) Volume() float64 {
	if s == nil {
		return 0
	}

	return (float64(4) / 3) * math.Pi * math.Pow(s.radius, 3)
}
