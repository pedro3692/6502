package cpu

const (
	brkCost = 7
	rtiCost = 6
)

func (cpu *CPU) brk() int {
	cpu.p.SetBreak()
	cpu.p.ResetDecimalMode()
	cpu.p.SetIRQBDisable()

	cpu.pc.Read()
	cpu.pc.Read()
	pc := cpu.pc.Read() // PC + 2

	cpu.memory.StackPush(cpu.sp.Inc(), pc[1])        // push PCH
	cpu.memory.StackPush(cpu.sp.Inc(), pc[0])        // push PCL
	cpu.memory.StackPush(cpu.sp.Inc(), cpu.p.Read()) // push P

	pcl := cpu.memory.Read([2]byte{0xfe, 0xff})
	pch := cpu.memory.Read([2]byte{0xff, 0xff})

	cpu.pc.Load([2]byte{pcl, pch})

	return brkCost
}

func (cpu *CPU) rti() int {
	cpu.pc.Read()

	cpu.p.Load(cpu.memory.StackPull(cpu.sp.Dec())) // pull P
	pcl := cpu.memory.StackPull(cpu.sp.Dec())      // pull PCL
	pch := cpu.memory.StackPull(cpu.sp.Dec())      // pull PCH

	cpu.pc.Load([2]byte{pcl, pch})

	return rtiCost
}
