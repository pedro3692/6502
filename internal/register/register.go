package register

type Memory interface {
	Store([2]byte, byte)
	Read([2]byte) byte
}

type Register struct {
	data   byte
	memory Memory
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

func New(mem Memory) Register {
	return Register{
		memory: mem,
	}
}
