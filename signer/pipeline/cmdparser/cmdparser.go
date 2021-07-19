package cmdparser

import (
	"fmt"
	"strings"
)

const operators = "|&>"

type cmdInformation struct {
	name              string
	args              []string
	pipelineOperation string
}

func parseCmd(cmd string, pipeOperation string) (cmdInformation, error) {
	partsOfCmd := strings.Fields(cmd)

	//TODO: create function for assert combinations
	if !strings.ContainsAny(pipeOperation, operators) {
		return cmdInformation{}, fmt.Errorf("operator %v isn't valid", pipeOperation)
	}

	return cmdInformation{
		name:              partsOfCmd[0],
		args:              partsOfCmd[1:len(partsOfCmd)],
		pipelineOperation: pipeOperation,
	}, nil
}

func (i cmdInformation) getName() string {
	return i.name
}

func (i cmdInformation) getArgs() string {
	return i.name + strings.Join(i.args, " ")
}

func (i cmdInformation) getPipelineOperation() string {
	return i.pipelineOperation
}

func Parse(input string) ([]cmdInformation, error) {
	cmds := make([]cmdInformation, 0, 2)

	for end := strings.IndexAny(input, operators); end != -1; end = strings.IndexAny(input, operators) {
		startNextCmd := end + strings.Index(input[end:], " ") + 1
		cmdInfo, err := parseCmd(input[0:end], input[end:startNextCmd-1])
		if err != nil {
			return nil, fmt.Errorf("parse input command: %v", err)
		}
		cmds = append(cmds, cmdInfo)

		input = input[startNextCmd:]
	}

	return cmds, nil
}
