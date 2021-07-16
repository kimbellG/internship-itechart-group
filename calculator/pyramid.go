package main

import (
	"flag"
	"internship-itechart-group/calculator/figure"
	"internship-itechart-group/calculator/figure/pyramid"
	"os"
)

var (
	pyramidCmd      = flag.NewFlagSet("pyramid", flag.ExitOnError)
	pyramidBaseArea = pyramidCmd.Float64("S", 0, "set the base area of pyramid")
	pyramidHeight   = pyramidCmd.Float64("h", 0, "set height of pyramid")
)

func pyramidVolume() figure.FigureVolume {
	pyramidCmd.Parse(os.Args[2:])

	return figure.FigureVolume{"pyramid", pyramid.NewPyramid(*pyramidBaseArea, *pyramidHeight)}

}
