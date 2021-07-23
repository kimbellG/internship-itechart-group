package hashout

import (
	"fmt"
	"sync"
)

type HashOutput struct {
	combineHash string

	sync.Mutex
}

func New() *HashOutput {
	return &HashOutput{}
}

func (h *HashOutput) CombineHash() string {
	if h.combineHash == "" {
		return ""
	}

	return h.combineHash[1:]
}

func (h *HashOutput) Print(name string, output []byte) {
	fmt.Printf("Output of %v:\n %s", name, output)
	hash := string(printHash(name, output))

	h.addPartOfCombineHash(hash)
}

func (h *HashOutput) addPartOfCombineHash(hash string) {
	h.Lock()
	defer h.Unlock()

	h.combineHash += "_" + hash
}
