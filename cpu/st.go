package cpu

import "github.com/pedro3692/6502/internal/register"

const (
	stAbsCost   = 4
	stZpCost    = 3
	stZpxCost   = 4
	stZpyCost   = 4
	stAbsxCost  = 5
	stAbsyCost  = 5
	stAIndxCost = 6
	stAIndyCost = 6
)

func (cpu *CPU) staAbs() int {
	return cpu.stAbs(&cpu.a)
}

func (cpu *CPU) stxAbs() int {
	return cpu.stAbs(&cpu.x)
}

func (cpu *CPU) styAbs() int {
	return cpu.stAbs(&cpu.y)
}

func (cpu *CPU) staZp() int {
	return cpu.stZp(&cpu.a)
}

func (cpu *CPU) stxZp() int {
	return cpu.stZp(&cpu.x)
}

func (cpu *CPU) styZp() int {
	return cpu.stZp(&cpu.y)
}

func (cpu *CPU) staZpx() int {
	return cpu.stZpx(&cpu.a)
}

func (cpu *CPU) staAbsx() int {
	cost := stAbsxCost

	baseLb := cpu.bus.Read(cpu.pc.Read())
	baseHb := cpu.bus.Read(cpu.pc.Read())

	x := cpu.x.Read()

	addr := int16(baseLb) | int16(baseHb)<<8 + int16(x)

	lb := byte(addr & 0xFF)
	hb := byte(addr >> 8)

	cpu.a.Store([2]byte{lb, hb})

	if int16(baseLb+x) >= 0x100 {
		cost++
	}

	return cost
}

func (cpu *CPU) staAbsy() int {
	cost := stAbsyCost

	baseLb := cpu.bus.Read(cpu.pc.Read())
	baseHb := cpu.bus.Read(cpu.pc.Read())

	y := cpu.y.Read()

	addr := int16(baseLb) | int16(baseHb)<<8 + int16(y)

	lb := byte(addr & 0xFF)
	hb := byte(addr >> 8)

	cpu.a.Store([2]byte{lb, hb})

	if int16(baseLb+y) >= 0x100 {
		cost++
	}

	return cost
}

func (cpu *CPU) staIndx() int {
	baseLb := cpu.bus.Read(cpu.pc.Read())

	lbSum := int16(baseLb) + int16(cpu.x.Read())

	lb := cpu.bus.Read([2]byte{byte(lbSum & 0xFF), 0x00})

	baseHb := byte((int16(lb) + int16(0x01)) & 0xFF)
	hb := cpu.bus.Read([2]byte{baseHb, 0x00})

	cpu.a.Store([2]byte{lb, hb})

	return stAIndxCost
}

func (cpu *CPU) staIndy() int {
	baseLb := cpu.bus.Read(cpu.pc.Read())
	baseHb := cpu.bus.Read([2]byte{baseLb + 0x01, 0x00})

	lbSum := int16(cpu.bus.Read([2]byte{baseLb, 0x00})) + int16(cpu.y.Read())

	lb := byte(lbSum & 0xFF)
	overflow := byte(lbSum >> 8)
	hb := baseHb + overflow

	cpu.a.Store([2]byte{lb, hb})

	return stAIndyCost
}

func (cpu *CPU) stxZpy() int {
	lb := byte(0xff)
	base := cpu.bus.Read(cpu.pc.Read())
	y := cpu.y.Read()

	if int16(base+y) < 0x100 {
		lb = base + y
	}

	cpu.x.Store([2]byte{lb, 0x00})

	return stZpyCost
}

func (cpu *CPU) styZpx() int {
	return cpu.stZpx(&cpu.y)
}

func (cpu *CPU) stAbs(r *register.Register) int {
	lb, hb := cpu.abs()

	r.Store([2]byte{lb, hb})

	return stAbsCost
}

func (cpu *CPU) stZp(r *register.Register) int {
	r.Store([2]byte{cpu.bus.Read(cpu.pc.Read()), 0x00})

	return stZpCost
}

func (cpu *CPU) stZpx(r *register.Register) int {
	lb := byte(0xff)
	base := cpu.bus.Read(cpu.pc.Read())
	x := cpu.x.Read()

	if int16(base+x) < 0x100 {
		lb = base + x
	}

	r.Store([2]byte{lb, 0x00})

	return stZpxCost
}
