package cpu

//go:generate stringer -type=Instruction
type Instruction byte

const (
	BRK       Instruction = 0x00
	CLC       Instruction = 0x18
	SEC       Instruction = 0x38
	RTI       Instruction = 0x40
	JMP_ABS   Instruction = 0x4c
	CLI       Instruction = 0x58
	ADC_IND_X Instruction = 0x61
	ADC_ZP    Instruction = 0x65
	ADC_IMM   Instruction = 0x69
	ADC_ABS   Instruction = 0x6D
	ADC_IND_Y Instruction = 0x71
	ADC_ZP_X  Instruction = 0x75
	SEI       Instruction = 0x78
	ADC_ABS_Y Instruction = 0x79
	ADC_ABS_X Instruction = 0x7D
	STA_IND_X Instruction = 0x81
	STY_ZP    Instruction = 0x84
	STA_ZP    Instruction = 0x85
	STX_ZP    Instruction = 0x86
	STY_ABS   Instruction = 0x8c
	STA_ABS   Instruction = 0x8d
	STX_ABS   Instruction = 0x8e
	STA_IND_Y Instruction = 0x91
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
	CLV       Instruction = 0xb8
	CLD       Instruction = 0xd8
	SBC_IND_X Instruction = 0xe1
	SBC_ZP    Instruction = 0xe5
	SBC_IMM   Instruction = 0xe9
	NOP       Instruction = 0xea
	SBC_ABS   Instruction = 0xed
	SBC_IND_Y Instruction = 0xf1
	SBC_ZP_X  Instruction = 0xf5
	SED       Instruction = 0xf8
	SBC_ABS_Y Instruction = 0xf9
	SBC_ABS_X Instruction = 0xfd
)

func (cpu *CPU) createInstuctionsTable() map[Instruction]instructionFunc {
	instructionTable := make(map[Instruction]instructionFunc, 256)

	instructionTable[BRK] = cpu.brk

	instructionTable[CLC] = cpu.clc

	instructionTable[RTI] = cpu.rti

	instructionTable[SEC] = cpu.sec

	instructionTable[JMP_ABS] = cpu.jmpAbs

	instructionTable[CLV] = cpu.clv
	instructionTable[CLI] = cpu.cli

	instructionTable[ADC_IND_X] = cpu.adcIndx
	instructionTable[ADC_ZP] = cpu.adcZp
	instructionTable[ADC_IMM] = cpu.adcImm
	instructionTable[ADC_ABS] = cpu.adcAbs
	instructionTable[ADC_IND_Y] = cpu.adcIndy
	instructionTable[ADC_ZP_X] = cpu.adcZpx

	instructionTable[SEI] = cpu.sei

	instructionTable[ADC_ABS_Y] = cpu.adcAbsy
	instructionTable[ADC_ABS_X] = cpu.adcAbsx
	instructionTable[ADC_ZP_X] = cpu.adcZpx

	instructionTable[STA_IND_X] = cpu.staIndx
	instructionTable[STY_ZP] = cpu.styZp
	instructionTable[STA_ZP] = cpu.staZp
	instructionTable[STX_ZP] = cpu.stxZp
	instructionTable[STY_ABS] = cpu.styAbs
	instructionTable[STA_ABS] = cpu.staAbs
	instructionTable[STX_ABS] = cpu.stxAbs
	instructionTable[STA_IND_Y] = cpu.staIndy
	instructionTable[STY_ZP_X] = cpu.styZpx
	instructionTable[STA_ZP_X] = cpu.staZpx
	instructionTable[STX_ZP_Y] = cpu.stxZpy
	instructionTable[STA_ABS_Y] = cpu.staAbsy
	instructionTable[STA_ABS_X] = cpu.staAbsx

	instructionTable[LDY_IMM] = cpu.ldyImm
	instructionTable[LDX_IMM] = cpu.ldxImm
	instructionTable[LDY_ZP] = cpu.ldyZp
	instructionTable[LDA_ZP] = cpu.ldaZp
	instructionTable[LDX_ZP] = cpu.ldxZp
	instructionTable[LDA_IMM] = cpu.ldaImm
	instructionTable[LDY_ABS] = cpu.ldyAbs
	instructionTable[LDA_ABS] = cpu.ldaAbs
	instructionTable[LDX_ABS] = cpu.ldxAbs

	instructionTable[CLD] = cpu.cld

	instructionTable[SBC_IND_X] = cpu.sbcIndx
	instructionTable[SBC_ZP] = cpu.sbcZp
	instructionTable[SBC_IMM] = cpu.sbcImm
	instructionTable[SBC_ABS] = cpu.sbcAbs
	instructionTable[SBC_IND_Y] = cpu.sbcIndy
	instructionTable[SBC_ZP_X] = cpu.sbcZpx

	instructionTable[SED] = cpu.sed

	instructionTable[SBC_ABS_Y] = cpu.sbcAbsy
	instructionTable[SBC_ABS_X] = cpu.sbcAbsx
	instructionTable[SBC_ZP_X] = cpu.sbcZpx

	instructionTable[NOP] = cpu.nop

	return instructionTable
}
