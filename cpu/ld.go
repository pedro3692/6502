package cpu

import "github.com/pedro3692/6502/internal/register"

const (
	ldImmCost = 2
	ldAbsCost = 4
	ldZpCost  = 3
)

func (cpu *CPU) ldaImm() int {
	return cpu.ldImm(&cpu.a)
}

func (cpu *CPU) ldxImm() int {
	return cpu.ldImm(&cpu.x)
}

func (cpu *CPU) ldyImm() int {
	return cpu.ldImm(&cpu.y)
}

func (cpu *CPU) ldaAbs() int {
	return cpu.ldAbs(&cpu.a)
}

func (cpu *CPU) ldxAbs() int {
	return cpu.ldAbs(&cpu.x)
}

func (cpu *CPU) ldyAbs() int {
	return cpu.ldAbs(&cpu.y)
}

func (cpu *CPU) ldaZp() int {
	return cpu.ldZp(&cpu.a)
}

func (cpu *CPU) ldxZp() int {
	return cpu.ldZp(&cpu.x)
}

func (cpu *CPU) ldyZp() int {
	return cpu.ldZp(&cpu.y)
}

func (cpu *CPU) ldImm(r *register.Register) int {
	cpu.p.ResetNegative()
	cpu.p.ResetZero()

	data := cpu.imm()
	r.Load(data)

	if data == 0x00 {
		cpu.p.SetZero()
	}

	if data>>7 == 1 {
		cpu.p.SetNegative()
	}

	return ldImmCost
}

func (cpu *CPU) ldAbs(r *register.Register) int {
	cpu.p.ResetNegative()
	cpu.p.ResetZero()

	lb, hb := cpu.abs()

	data := cpu.memory.Read([2]byte{lb, hb})

	r.Load(data)

	if data == 0x00 {
		cpu.p.SetZero()
	}

	if data>>7 == 1 {
		cpu.p.SetNegative()
	}

	return ldAbsCost
}

func (cpu *CPU) ldZp(r *register.Register) int {
	cpu.p.ResetNegative()
	cpu.p.ResetZero()

	data := cpu.zp()

	r.Load(data)

	if data == 0x00 {
		cpu.p.SetZero()
	}

	if data>>7 == 1 {
		cpu.p.SetNegative()
	}

	return ldZpCost
}
