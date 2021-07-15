package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()

	for i, str := range flag.Args() {
		fmt.Printf("%v: %v\n", i+1, str)
		for k, v := range getWordCountOfString(str) {
			fmt.Printf("\t%v: %v", k, v)
		}
		fmt.Println()
	}

}

func getWordCountOfString(str string) map[string]int {
	result := make(map[string]int)
	for _, word := range strings.Fields(str) {
		result[word]++
	}

	return result
}
