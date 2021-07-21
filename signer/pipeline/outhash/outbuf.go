package outhash

import (
	"fmt"
	"sync"
)

type outputBuffer struct {
	sync.Mutex
	result          string
	currentHash     int
	single          string
	indexOfProgramm int
	buffer          []multiHash
}

func (o *outputBuffer) printHash(hash multiHash) {
	o.Lock()
	defer o.Unlock()

	if hash.indexOfInteration == o.currentHash {
		o.printMultiHash(hash.indexOfInteration, hash.hash)
	} else {
		o.buffer = append(o.buffer, hash)
		return
	}

	o.clearBuffer()

}

func (o *outputBuffer) printMultiHash(index int, hash uint32) {
	fmt.Printf("%d MultuHash %d %s %d\n", o.indexOfProgramm, index, o.single, hash)
	o.result = fmt.Sprintf("%s%d", o.result, hash)
	o.currentHash++
}

func (o *outputBuffer) clearBuffer() {
	for isCurrentHashInBuffer := true; isCurrentHashInBuffer; {
		isCurrentHashInBuffer = false
		for i, h := range o.buffer {
			if h.indexOfInteration == o.currentHash {
				o.printMultiHash(h.indexOfInteration, h.hash)
				o.delFromBuffer(i)
				isCurrentHashInBuffer = true
				break
			}
		}
	}
}

func (o *outputBuffer) delFromBuffer(i int) {
	if len(o.buffer) == 1 {
		o.buffer = []multiHash{}
		return
	}

	copy(o.buffer[i:], o.buffer[i+1:])
	o.buffer = o.buffer[:len(o.buffer)-1]
}
