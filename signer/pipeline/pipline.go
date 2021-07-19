package pipeline

import (
	"fmt"
	"internship-itechart-group/signer/pipeline/cmdparser"
)

func Execute(input string) (string, error) {
	cmds, err := cmdparser.Parse(input)
	if err != nil {
		return "", fmt.Errorf("parse commands: %v", err)
	}

}
