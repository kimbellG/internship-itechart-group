package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()

	if err := printTree(flag.Args()); err != nil {
		fmt.Printf("word-count: %v\n", err)
	}
}

func printTree(strs []string) error {
	if len(strs) == 0 {
		return fmt.Errorf("expected at least one string of words as an argument")
	}

	for i, str := range strs {
		fmt.Printf("%v: %v\n", i+1, str)
		for k, v := range getWordCountOfString(str) {
			fmt.Printf("\t%v: %v", k, v)
		}
		fmt.Println()
	}

	return nil
}

func getWordCountOfString(str string) map[string]int {
	result := make(map[string]int)
	for _, word := range strings.Fields(str) {
		result[word]++
	}

	return result
}
