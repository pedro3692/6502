package register

type ProgramCounter struct {
	data [2]byte
}

func (pc *ProgramCounter) Load(b [2]byte) {
	pc.data = b
}

func (pc *ProgramCounter) Read() [2]byte {
	data := pc.data

	pc.inc()

	return data
}

func (pc *ProgramCounter) Check() [2]byte {
	return pc.data
}

func (pc *ProgramCounter) Reset() {
	pc.data = [2]byte{0xfc, 0xff}
}

func (pc *ProgramCounter) inc() {
	if pc.data[0] == 0xff {
		pc.data[0] = 0x00
		pc.data[1]++

		return
	}

	pc.data[0]++
}
