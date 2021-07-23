package cmdparser

import (
	"fmt"
	"strings"
)

const operators = "|"

type CMDInformation struct {
	name              string
	args              []string
	pipelineOperation string
}

func parseCmd(cmd string, pipeOperation string) (CMDInformation, error) {
	partsOfCmd := strings.Fields(cmd)

	if cmd == "" {
		return CMDInformation{}, fmt.Errorf("command is not defined")
	}

	//TODO: create function for assert combinations
	if !strings.ContainsAny(pipeOperation, operators) && pipeOperation != "" {
		return CMDInformation{}, fmt.Errorf("operator %v isn't valid", pipeOperation)
	}

	return CMDInformation{
		name:              partsOfCmd[0],
		args:              partsOfCmd[1:len(partsOfCmd)],
		pipelineOperation: pipeOperation,
	}, nil
}

func (i CMDInformation) Name() string {
	return i.name
}

func (i CMDInformation) Args() []string {
	return i.args
}

func (i CMDInformation) PipelineOperation() string {
	return i.pipelineOperation
}

func Parse(input string) ([]CMDInformation, error) {
	cmds := make([]CMDInformation, 0, 2)

	for end := strings.IndexAny(input, operators); end != -1; end = strings.IndexAny(input, operators) {
		startNextCmd := end + strings.Index(input[end:], " ") + 1
		cmdInfo, err := parseCmd(input[0:end], input[end:startNextCmd-1])
		if err != nil {
			return nil, fmt.Errorf("parse input command: %v", err)
		}
		cmds = append(cmds, cmdInfo)

		input = input[startNextCmd:]
	}

	lastCmd, err := parseCmd(input[0:], "")
	if err != nil {
		return nil, fmt.Errorf("parse input command: %v", err)
	}

	return append(cmds, lastCmd), nil
}
