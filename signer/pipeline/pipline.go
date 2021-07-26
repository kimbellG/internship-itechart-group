package pipeline

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"

	"github.com/kimbellG/internship-itechart-group/signer/pipeline/cmdparser"
)

type Pipeline struct {
	commandsInformation []cmdparser.CMDInformation
	pr                  Printer

	sync.Mutex
	wg sync.WaitGroup
}

type Printer interface {
	Print(name string, output []byte)
}

func NewPipeline(input string, pr Printer) (*Pipeline, error) {
	cmdInfos, err := cmdparser.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("parse commands: %v", err)
	}

	return &Pipeline{
		commandsInformation: cmdInfos,
		pr:                  pr,
	}, nil

}

func (p *Pipeline) Execute() error {
	var cmdOutput = []byte{}
	var err error

	for _, cmdInfo := range p.commandsInformation {
		cmdOutput, err = execute(cmdInfo.Name(), cmdOutput, cmdInfo.Args()...)
		if err != nil {
			return fmt.Errorf("execute %v: %v", cmdInfo.Name(), err)
		}

		p.wg.Add(1)
		go func(nameOfProgramm string, output []byte) {
			defer p.wg.Done()
			p.pr.Print(nameOfProgramm, output)

		}(cmdInfo.Name(), getCopyOfOutput(cmdOutput))
	}

	p.wg.Wait()

	return nil
}

func execute(name string, stdinData []byte, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)

	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("stdin pipe: %v", err)
	}

	go func() {
		defer closeWriter(in)

		if _, err := in.Write(stdinData); err != nil {
			log.Printf("write in stdin of %v: %v", name, err)
		}
	}()

	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("execute %v:\n %s\n%v", name, cmdOutput, err)
	}

	return cmdOutput, nil
}

func closeWriter(in io.WriteCloser) {
	if err := in.Close(); err != nil {
		log.Printf("error: close stdin pipe: %v", err)
	}
}

func getCopyOfOutput(output []byte) []byte {
	cpOfOutput := make([]byte, len(output))
	copy(cpOfOutput, output)
	return cpOfOutput
}
