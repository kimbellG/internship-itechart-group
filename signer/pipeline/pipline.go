package pipeline

import (
	"fmt"
	"internship-itechart-group/signer/pipeline/cmdparser"
	"internship-itechart-group/signer/pipeline/outhash"
	"io"
	"log"
	"os/exec"
	"sync"
)

type Pipeline struct {
	commandsInformation []cmdparser.CMDInformation
	combineHash         string
	funcForOutput       func(nameOfProgramm string, output []byte)

	sync.Mutex
	wg sync.WaitGroup
}

func NewPipeline(input string, funcForOutput func(nameOfProgramm string, output []byte)) (*Pipeline, error) {
	cmdInfos, err := cmdparser.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("parse commands: %v", err)
	}

	return &Pipeline{
		commandsInformation: cmdInfos,
		funcForOutput:       funcForOutput,
	}, nil

}

func (p *Pipeline) CombineHash() string {
	if p.combineHash == "" {
		return ""
	}

	return p.combineHash[1:]
}

func (p *Pipeline) Execute() error {
	cmdOutput, err := []byte{}, error(nil)

	for i, cmdInfo := range p.commandsInformation {
		cmdOutput, err = execute(cmdInfo.Name(), cmdOutput, cmdInfo.Args()...)
		if err != nil {
			return fmt.Errorf("execute %v: %v", cmdInfo.Name(), err)
		}
		log.Printf("Output of %v:\n %s", cmdInfo.Name(), cmdOutput)

		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			p.PrintHashOfOutput(i, getCopyOfOutput(cmdOutput))
		}()
	}

	p.wg.Wait()

	fmt.Printf("Combine Result: %v\n", p.CombineHash())
	return nil
}

func (p *Pipeline) PrintHashOfOutput(indexOfProgramm int, output []byte) {
	hash := string(outhash.PrintHash(indexOfProgramm, output))

	p.addPartOfCombineHash(hash)
}

func (p *Pipeline) addPartOfCombineHash(hash string) {
	p.Lock()
	defer p.Unlock()

	p.combineHash += "_" + hash
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
	cpOfOutput := []byte{}
	copy(cpOfOutput, output)
	return cpOfOutput
}
