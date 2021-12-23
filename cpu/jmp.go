package cpu

const jmpAbsCost = 3

func (cpu *CPU) jmpAbs() int {
	pcl := cpu.bus.Read(cpu.pc.Read())
	pch := cpu.bus.Read(cpu.pc.Read())

	cpu.pc.Load([2]byte{pcl, pch})

	return jmpAbsCost
}
