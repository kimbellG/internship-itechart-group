package main

import (
	"flag"
	"internship-itechart-group/calculator/figure"
	"internship-itechart-group/calculator/figure/sphere"
	"os"
)

var (
	sphereCmd    = flag.NewFlagSet("sphere", flag.ExitOnError)
	sphereRadius = sphereCmd.Float64("r", 0, "set the radius of sphere")
)

func sphereVolume() figure.FigureVolume {
	sphereCmd.Parse(os.Args[2:])

	return figure.FigureVolume{"sphere", sphere.NewSphere(*sphereRadius)}
}
