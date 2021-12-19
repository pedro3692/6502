package memory

const (
	Size     = 65536
	Page     = 256
	ZeroPage = 0x0
	Stack    = 0x100
)

type Memory struct {
	memory [Size]byte
}

func (m *Memory) Store(address [2]byte, b byte) {
	littleEndianAdress := LittleEndian(address)

	m.memory[littleEndianAdress] = b
}

func (m Memory) Read(address [2]byte) byte {
	littleEndianAdress := LittleEndian(address)

	return m.memory[littleEndianAdress]
}

//TODO: remove this method after mapping the ROM
func (m *Memory) Set(memory [Size]byte) {
	m.memory = memory
}
