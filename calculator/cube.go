package main

import (
	"flag"
	"os"

	"internship-itechart-group/calculator/figure"
	"internship-itechart-group/calculator/figure/cube"
)

var (
	cubeCmd  = flag.NewFlagSet("cube", flag.ExitOnError)
	cubeSide = cubeCmd.Float64("s", 0, "set a side of cube")
)

func cubeVolume() figure.FigureVolume {
	cubeCmd.Parse(os.Args[2:])

	return figure.FigureVolume{"cube", cube.NewCube(*cubeSide)}
}
