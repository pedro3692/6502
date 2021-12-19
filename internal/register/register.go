package register

import (
	"github.com/pedro3692/6502/internal/memory"
)

type Register struct {
	data   byte
	memory *memory.Memory
}

func (r *Register) Load(b byte) {
	r.data = b
}

func (r *Register) Store(address [2]byte) {
	r.memory.Store(address, r.data)
}

func (r Register) Read() byte {
	return r.data
}

func New(mem *memory.Memory) Register {
	return Register{
		memory: mem,
	}
}
