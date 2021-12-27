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
