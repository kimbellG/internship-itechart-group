package main

import (
	"fmt"
	"internship-itechart-group/signer/pipeline"
	"os"
	"strings"
)

func main() {
	pipe, err := pipeline.NewPipeline(strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Printf("signer: %v\n", err)
		os.Exit(1)
	}

	if err := pipe.Execute(strings.Join(os.Args[1:], " ")); err != nil {
		fmt.Printf("signer: %v\n", err)
		os.Exit(1)
	}
}
