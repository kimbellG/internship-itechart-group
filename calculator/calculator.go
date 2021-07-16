package main

import (
	"fmt"
	"log"
	"os"

	"internship-itechart-group/calculator/figure"
)

func main() {
	if err := checkInputFigure(); err != nil {
		log.Fatalf("calculator: %v\n", err)
	}

	volume, err := getVolumeForFigure(os.Args[1])
	if err != nil {
		fmt.Printf("calculator: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(volume)
}

func checkInputFigure() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("expected figure for calculator")
	}

	return nil
}

func getVolumeForFigure(fig string) (figure.FigureVolume, error) {
	switch os.Args[1] {

	case "cube":
		return cubeVolume(), nil
	case "sphere":
		return sphereVolume(), nil
	case "pyramid":
		return pyramidVolume(), nil
	default:
		return figure.FigureVolume{}, fmt.Errorf("%v isn't supported", fig)
	}
}
