package cpu

const (
	adcImmCost  = 2
	adcZPCost   = 3
	adcZpxCost  = 4
	adcAbsCost  = 4
	adcAbsxCost = 4
	adcAbsyCost = 4
	adcIndxCost = 6
	adcIndyCost = 5
)

func (cpu *CPU) adcImm() int {
	cpu.adc(cpu.imm())

	return adcImmCost
}

func (cpu *CPU) adcAbs() int {
	cpu.adc(cpu.abs())

	return adcAbsCost
}

func (cpu *CPU) adcAbsx() int {
	cost := adcAbsCost
	data, pageCrossed := cpu.absx()

	if pageCrossed {
		cost++
	}

	cpu.adc(data)

	return cost
}

func (cpu *CPU) adcAbsy() int {
	cost := adcAbsCost
	data, pageCrossed := cpu.absy()

	if pageCrossed {
		cost++
	}

	cpu.adc(data)

	return cost
}

func (cpu *CPU) adcZp() int {
	cpu.adc(cpu.zp())

	return adcZPCost
}

func (cpu *CPU) adcZpx() int {
	cpu.adc(cpu.zpx())

	return adcZpxCost
}

func (cpu *CPU) adcIndx() int {
	cpu.adc(cpu.indx())

	return adcIndxCost
}

func (cpu *CPU) adcIndy() int {
	cost := adcIndyCost
	data, pageCrossed := cpu.indy()

	if pageCrossed {
		cost++
	}

	cpu.adc(data)

	return cost
}

func (cpu *CPU) adc(data byte) {
	cpu.p.ResetZero()
	cpu.p.ResetNegative()

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

	byteSum := byte(sum & 0xFF)

	if byteSum>>7 == 1 {
		cpu.p.SetNegative()
	}

	if byteSum == 0x00 {
		cpu.p.SetZero()
	}

	cpu.a.Load(byteSum)
}
