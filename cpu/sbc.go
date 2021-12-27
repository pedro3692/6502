package cpu

const (
	sbcImmCost  = 2
	sbcZpCost   = 3
	sbcZpxCost  = 4
	sbcAbsCost  = 4
	sbcAbsxCost = 4
	sbcAbsyCost = 4
	sbcIndxCost = 6
	sbcIndyCost = 5
)

func (cpu *CPU) sbcImm() int {
	cpu.sbc(cpu.imm())

	return sbcImmCost
}

func (cpu *CPU) sbcAbs() int {
	cpu.sbc(cpu.abs())

	return sbcAbsCost
}

func (cpu *CPU) sbcAbsx() int {
	cost := sbcAbsCost
	data, pageCrossed := cpu.absx()

	if pageCrossed {
		cost++
	}

	cpu.sbc(data)

	return cost
}

func (cpu *CPU) sbcAbsy() int {
	cost := sbcAbsCost
	data, pageCrossed := cpu.absy()

	if pageCrossed {
		cost++
	}

	cpu.sbc(data)

	return cost
}

func (cpu *CPU) sbcZp() int {
	cpu.sbc(cpu.zp())

	return sbcZpCost
}

func (cpu *CPU) sbcZpx() int {
	cpu.sbc(cpu.zpx())

	return sbcZpxCost
}

func (cpu *CPU) sbcIndx() int {
	cpu.sbc(cpu.indx())

	return sbcIndxCost
}

func (cpu *CPU) sbcIndy() int {
	cost := sbcIndyCost
	data, pageCrossed := cpu.indy()

	if pageCrossed {
		cost++
	}

	cpu.sbc(data)

	return cost
}

func (cpu *CPU) sbc(data byte) {
	cpu.p.ResetZero()
	cpu.p.ResetNegative()

	a := cpu.a.Read()
	carry := cpu.p.Carry()

	data = ^data

	if carry {
		data = data + 0x1
	}

	sub := uint16(data) + uint16(a)

	cpu.p.ResetCarry()

	if sub>>8 != 0x00 {
		cpu.p.SetCarry()
	}

	byteSum := byte(sub & 0xFF)

	if byteSum>>7 == 1 {
		cpu.p.SetNegative()
	}

	if byteSum == 0x00 {
		cpu.p.SetZero()
	}

	cpu.a.Load(byteSum)
}
