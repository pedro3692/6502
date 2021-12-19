package cpu

import "github.com/pedro3692/6502/internal/register"

const ldImCost = 2

func (cpu *CPU) ldaIm() int {
	return cpu.ldIm(&cpu.a)
}

func (cpu *CPU) ldxIm() int {
	return cpu.ldIm(&cpu.x)
}

func (cpu *CPU) ldyIm() int {
	return cpu.ldIm(&cpu.y)
}

func (cpu *CPU) ldIm(r *register.Register) int {
	cpu.p.ResetNegative()
	cpu.p.ResetZero()

	data := cpu.memory.Read(cpu.pc.Read())
	r.Load(data)

	if data == 0x00 {
		cpu.p.SetZero()
	}

	if data>>7 == 1 {
		cpu.p.SetNegative()
	}

	return ldImCost
}
