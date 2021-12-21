package cpu

func (cpu *CPU) imm() byte {
	return cpu.memory.Read(cpu.pc.Read())
}

func (cpu *CPU) abs() (lb, hb byte) {
	lb = cpu.memory.Read(cpu.pc.Read())
	hb = cpu.memory.Read(cpu.pc.Read())

	return
}

func (cpu *CPU) zp() byte {
	lb := cpu.memory.Read(cpu.pc.Read())

	return cpu.memory.Read([2]byte{lb, 0x00})
}

func (cpu *CPU) zpx() byte {
	lb := byte(0xff)
	base := cpu.memory.Read(cpu.pc.Read())
	x := cpu.x.Read()

	if int16(base+x) < 0x100 {
		lb = base + x
	}

	return cpu.memory.Read([2]byte{lb, 0x00})
}
