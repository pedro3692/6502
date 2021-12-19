package cpu

import "github.com/pedro3692/6502/internal/register"

const stAbsCost = 4

func (cpu *CPU) staAbs() int {
	return cpu.stAbs(&cpu.a)
}

func (cpu *CPU) stxAbs() int {
	return cpu.stAbs(&cpu.x)
}

func (cpu *CPU) styAbs() int {
	return cpu.stAbs(&cpu.y)
}

func (cpu *CPU) stAbs(r *register.Register) int {
	r.Store([2]byte{cpu.memory.Read(cpu.pc.Read()), cpu.memory.Read(cpu.pc.Read())})

	return stAbsCost
}
