package cpu

var (
	clearCost = 2
	setCost   = 2
)

func (cpu *CPU) sec() int {
	cpu.p.SetCarry()

	return setCost
}

func (cpu *CPU) sed() int {
	cpu.p.SetDecimalMode()

	return setCost
}

func (cpu *CPU) sei() int {
	cpu.p.SetIRQBDisable()

	return setCost
}

func (cpu *CPU) clc() int {
	cpu.p.ResetCarry()

	return clearCost
}

func (cpu *CPU) cld() int {
	cpu.p.ResetDecimalMode()

	return clearCost
}

func (cpu *CPU) cli() int {
	cpu.p.ResetIRQBDisable()

	return clearCost
}

func (cpu *CPU) clv() int {
	cpu.p.ResetOverflow()

	return clearCost
}
