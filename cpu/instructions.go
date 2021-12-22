package cpu

//go:generate stringer -type=Instruction
type Instruction byte

const (
	BRK       Instruction = 0x00
	RTI       Instruction = 0x40
	JMP_ABS   Instruction = 0x4c
	STY_ZP    Instruction = 0x84
	STA_ZP    Instruction = 0x85
	STX_ZP    Instruction = 0x86
	STY_ABS   Instruction = 0x8c
	STA_ABS   Instruction = 0x8d
	STX_ABS   Instruction = 0x8e
	STA_ZP_Y  Instruction = 0x91
	STY_ZP_X  Instruction = 0x94
	STA_ZP_X  Instruction = 0x95
	STX_ZP_Y  Instruction = 0x96
	STA_ABS_Y Instruction = 0x99
	STA_ABS_X Instruction = 0x9D
	LDY_IMM   Instruction = 0xa0
	LDX_IMM   Instruction = 0xa2
	LDY_ZP    Instruction = 0xa4
	LDA_ZP    Instruction = 0xa5
	LDX_ZP    Instruction = 0xa6
	LDA_IMM   Instruction = 0xa9
	LDY_ABS   Instruction = 0xac
	LDA_ABS   Instruction = 0xad
	LDX_ABS   Instruction = 0xae
)
