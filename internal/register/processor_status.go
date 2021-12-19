package register

type ProcessorStatus struct {
	data byte
}

func (ps ProcessorStatus) Negative() bool {
	return (ps.data&0b10000000)>>7 == 1
}

func (ps ProcessorStatus) Overflow() bool {
	return (ps.data&0b01000000)>>6 == 1
}

func (ps ProcessorStatus) Break() bool {
	return (ps.data&0b00010000)>>4 == 1
}

func (ps ProcessorStatus) DecimalMode() bool {
	return (ps.data&0b00001000)>>3 == 1
}

func (ps ProcessorStatus) IRQBDisable() bool {
	return (ps.data&0b00000100)>>2 == 1
}

func (ps ProcessorStatus) Zero() bool {
	return (ps.data&0b00000010)>>1 == 1
}

func (ps ProcessorStatus) Carry() bool {
	return (ps.data & 0b00000001) == 1
}

func (ps *ProcessorStatus) SetNegative() {
	ps.data = ps.data | 0b10000000
}

func (ps *ProcessorStatus) SetOverflow() {
	ps.data = ps.data | 0b01000000
}

func (ps *ProcessorStatus) SetBreak() {
	ps.data = ps.data | 0b00010000
}

func (ps *ProcessorStatus) SetDecimalMode() {
	ps.data = ps.data | 0b00001000
}

func (ps *ProcessorStatus) SetIRQBDisable() {
	ps.data = ps.data | 0b00000100
}

func (ps *ProcessorStatus) SetZero() {
	ps.data = ps.data | 0b00000010
}

func (ps *ProcessorStatus) SetCarry() {
	ps.data = ps.data | 0b00000001
}

func (ps *ProcessorStatus) ResetNegative() {
	ps.data = ps.data & 0b01111111
}

func (ps *ProcessorStatus) ResetOverflow() {
	ps.data = ps.data & 0b10111111
}

func (ps *ProcessorStatus) ResetBreak() {
	ps.data = ps.data & 0b11101111
}

func (ps *ProcessorStatus) ResetDecimalMode() {
	ps.data = ps.data & 0b11110111
}

func (ps *ProcessorStatus) ResetIRQBDisable() {
	ps.data = ps.data & 0b11111011
}

func (ps *ProcessorStatus) ResetZero() {
	ps.data = ps.data & 0b11111101
}

func (ps *ProcessorStatus) ResetCarry() {
	ps.data = ps.data & 0b11111110
}
