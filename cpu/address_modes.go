package cpu

func (cpu *CPU) imm() byte {
	return cpu.bus.Read(cpu.pc.Read())
}

func (cpu *CPU) abs() byte {
	lb := cpu.bus.Read(cpu.pc.Read())
	hb := cpu.bus.Read(cpu.pc.Read())

	return cpu.bus.Read([2]byte{lb, hb})
}

func (cpu *CPU) absx() (byte, bool) {
	baseLb := cpu.bus.Read(cpu.pc.Read())
	hb := cpu.bus.Read(cpu.pc.Read())

	x := cpu.x.Read()

	lb := uint16(baseLb) + uint16(x)

	addr := lb + (uint16(hb) << 8)

	return cpu.bus.Read([2]byte{byte(addr & 0xFF), byte(addr >> 8)}), (lb >> 8) > 0x00
}

func (cpu *CPU) absy() (byte, bool) {
	baseLb := cpu.bus.Read(cpu.pc.Read())
	hb := cpu.bus.Read(cpu.pc.Read())

	y := cpu.y.Read()

	lb := uint16(baseLb) + uint16(y)

	addr := lb + (uint16(hb) << 8)

	return cpu.bus.Read([2]byte{byte(addr & 0xFF), byte(addr >> 8)}), (lb >> 8) > 0x00
}

func (cpu *CPU) zp() byte {
	lb := cpu.bus.Read(cpu.pc.Read())

	return cpu.bus.Read([2]byte{lb, 0x00})
}

func (cpu *CPU) zpx() byte {
	baseLb := cpu.bus.Read(cpu.pc.Read())
	x := cpu.x.Read()

	lb := byte((uint16(baseLb) + uint16(x)) & 0xff)

	return cpu.bus.Read([2]byte{lb, 0x00})
}

func (cpu *CPU) indx() byte {
	baseLb := cpu.bus.Read(cpu.pc.Read())
	x := cpu.x.Read()

	baseLb = byte((uint16(baseLb) + uint16(x)) & 0xff)

	lb := cpu.bus.Read([2]byte{baseLb, 0x00})
	hb := cpu.bus.Read([2]byte{baseLb + 0x01, 0x00})

	return cpu.bus.Read([2]byte{lb, hb})
}

func (cpu *CPU) indy() (byte, bool) {
	baseLb := cpu.bus.Read(cpu.pc.Read())
	baseHb := cpu.bus.Read([2]byte{baseLb + 0x01, 0x00})
	y := cpu.y.Read()

	lbSum := uint16(cpu.bus.Read([2]byte{baseLb, 0x00})) + uint16(y)

	lb := byte(lbSum & 0xff)
	overflow := byte(lbSum >> 8)
	hb := baseHb + overflow

	return cpu.bus.Read([2]byte{lb, hb}), overflow != 0x00
}
