package main

import (
	"fmt"
	"internship-itechart-group/signer/pipeline"
	"os"
	"strings"
)

func main() {
	output := pipeline.Execute(strings.Join(os.Args[1:], " "))

	fmt.Println(output)
}
