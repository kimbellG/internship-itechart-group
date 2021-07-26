package main

import (
	"flag"
	"fmt"
	"internship-itechart-group/signer/hashout"
	"internship-itechart-group/signer/pipeline"
	"os"
	"strings"
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
