package cpu

import "github.com/pedro3692/6502/internal/register"

const (
	cmpImmCost  = 2
	cmpZpCost   = 3
	cmpZpxCost  = 4
	cmpAbsCost  = 4
	cmpAbsxCost = 4
	cmpAbsyCost = 4
	cmpIndxCost = 6
	cmpIndyCost = 5
)

func (cpu *CPU) cmpImm() int {
	data := cpu.imm()

	cpu.cmp(&cpu.a, data)

	return cmpImmCost
}

func (cpu *CPU) cmpZp() int {
	data := cpu.zp()

	cpu.cmp(&cpu.a, data)

	return cmpZpCost
}

func (cpu *CPU) cmpZpx() int {
	data := cpu.zp()

	cpu.cmp(&cpu.a, data)

	return cmpZpxCost
}

func (cpu *CPU) cmpAbs() int {
	data := cpu.abs()

	cpu.cmp(&cpu.a, data)

	return cmpAbsCost
}

func (cpu *CPU) cmpAbsx() int {
	cost := cmpAbsxCost
	data, pageCrossed := cpu.absx()

	if pageCrossed {
		cost++
	}

	cpu.cmp(&cpu.a, data)

	return cost
}

func (cpu *CPU) cmpAbsy() int {
	cost := cmpAbsyCost
	data, pageCrossed := cpu.absy()

	if pageCrossed {
		cost++
	}

	cpu.cmp(&cpu.a, data)

	return cost
}

func (cpu *CPU) cmpIndx() int {
	data := cpu.indx()

	cpu.cmp(&cpu.a, data)

	return cmpAbsxCost
}

func (cpu *CPU) cmpIndy() int {
	cost := cmpIndyCost
	data, pageCrossed := cpu.indy()

	if pageCrossed {
		cost++
	}

	cpu.cmp(&cpu.a, data)

	return cost
}

func (cpu *CPU) cpxImm() int {
	data := cpu.imm()

	cpu.cmp(&cpu.x, data)

	return cmpImmCost
}

func (cpu *CPU) cpxZp() int {
	data := cpu.zp()

	cpu.cmp(&cpu.x, data)

	return cmpZpCost
}

func (cpu *CPU) cpxAbs() int {
	data := cpu.abs()

	cpu.cmp(&cpu.x, data)

	return cmpAbsCost
}

func (cpu *CPU) cpyImm() int {
	data := cpu.imm()

	cpu.cmp(&cpu.y, data)

	return cmpImmCost
}

func (cpu *CPU) cpyZp() int {
	data := cpu.zp()

	cpu.cmp(&cpu.y, data)

	return cmpZpCost
}

func (cpu *CPU) cpyAbs() int {
	data := cpu.abs()

	cpu.cmp(&cpu.y, data)

	return cmpAbsCost
}

func (cpu *CPU) cmp(r *register.Register, data byte) {
	cpu.p.ResetZero()
	cpu.p.ResetCarry()
	cpu.p.ResetNegative()

	rData := r.Read()

	result := rData - data

	if result>>7 == 1 {
		cpu.p.SetNegative()
	}

	if rData == data {
		cpu.p.SetZero()
	}

	if rData >= data {
		cpu.p.SetCarry()
	}
}
