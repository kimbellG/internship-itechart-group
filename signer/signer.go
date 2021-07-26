package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/internship-itechart-group/signer/hashout"
	"github.com/internship-itechart-group/signer/pipeline"
)

func main() {
	flag.Parse()

	hashPrinter := hashout.New()
	pipe, err := pipeline.NewPipeline(strings.Join(flag.Args(), " "), hashPrinter)
	if err != nil {
		fmt.Printf("signer: %v\n", err)
		os.Exit(1)
	}

	if err := pipe.Execute(); err != nil {
		fmt.Printf("signer: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Combined hash: %v", hashPrinter.CombineHash())
}
