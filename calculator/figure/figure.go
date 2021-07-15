package figure

import "fmt"

type FigureVolume struct {
	Name   string
	Volume Volumer
}

type Volumer interface {
	Volume() float64
}

func (f FigureVolume) String() string {
	return fmt.Sprintf("Volume of %v = %v\n", f.Name, f.Volume.Volume())
}
