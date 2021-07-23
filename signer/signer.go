package main

import (
	"flag"
	"fmt"
	"internship-itechart-group/signer/pipeline"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	pipe, err := pipeline.NewPipeline(strings.Join(flag.Args(), " "))
	if err != nil {
		fmt.Printf("signer: %v\n", err)
		os.Exit(1)
	}

	if err := pipe.Execute(); err != nil {
		fmt.Printf("signer: %v\n", err)
		os.Exit(1)
	}
}
