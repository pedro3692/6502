package memory

func (m Memory) Dump(from uint16, to uint16) []byte {
	return m.memory[from : to+1]
}
