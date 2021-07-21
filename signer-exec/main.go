package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	echoCmd := exec.Command("ipconfig")

	output, err := echoCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("comb out: %v", err)
		os.Exit(1)
	}

	printHash(output)

}
