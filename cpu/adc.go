package cpu

import "github.com/pedro3692/6502/internal/register"

const (
	adcImmCost = 2
	adcZPCost  = 3
	adcZPXCost = 4
	adcAbs     = 4
	adcAbsx    = 4
	adcAbsy    = 4
	adcIndx    = 6
	adcIndy    = 5
)

func (cpu *CPU) adcImm(r *register.Register) int {
	data := cpu.bus.Read(cpu.pc.Read())
	a := cpu.a.Read()
	carry := cpu.p.Carry()

	sum := uint16(data) + uint16(a)

	if carry {
		sum = sum + 0x1
	}

	cpu.p.ResetCarry()

	if sum>>8 != 0x00 {
		cpu.p.SetCarry()
	}

	cpu.a.Load(byte(sum & 0xFF))

	return adcImmCost
}
