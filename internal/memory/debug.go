package memory

func (m Memory) DumpStack() []byte {
	return m.memory[0x100 : 0x1ff+1]
}

func (m Memory) DumpZeroPage() []byte {
	return m.memory[0x000 : 0x0ff+1]
}

func (m Memory) Dump() []byte {
	return m.memory[:]
}
