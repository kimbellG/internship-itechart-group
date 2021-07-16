package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	startDirectory = flag.String("d", ".", "set the start directory")
)

func main() {
	flag.Parse()

	p, err := filepath.Abs(*startDirectory)
	if err != nil {
		log.Fatalf("tree: %v", err)
	}

	fmt.Printf("%v\n", p)
	if err := printDirectoryInformation(p, 0); err != nil {
		log.Fatalf("tree: %v", err)
	}
}

func printDirectoryInformation(dirpath string, hierarchyLevel int) error {
	files, err := os.ReadDir(dirpath)
	if err != nil {
		return fmt.Errorf("reading directory has been failed: %v", err)
	}

	for _, file := range files {
		fmt.Printf("%v---/%v\n", strings.Repeat("----", hierarchyLevel), file.Name())
		if file.IsDir() {
			if err := printDirectoryInformation(filepath.Join(dirpath, file.Name()), hierarchyLevel+1); err != nil {
				return err
			}
		}
	}

	return nil
}
