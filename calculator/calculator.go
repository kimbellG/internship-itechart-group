package main

import (
	"fmt"
	"log"
	"os"

	"internship-itechart-group/calculator/figure"
)

func main() {
	if err := checkInputFigure(); err != nil {
		log.Fatalf("calculator: %v", err)
	}

	var volume figure.FigureVolume
	switch os.Args[1] {

	case "cube":
		volume = cubeVolume()
	case "sphere":
		volume = sphereVolume()
	case "pyramid":
		volume = pyramidVolume()
	}

	fmt.Println(volume)
}

func checkInputFigure() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("expected figure for calculator")
	}

	return nil
}
