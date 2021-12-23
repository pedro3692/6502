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

	cpu.memory.StackPush(cpu.sp.Read(), pc[1]) // push PCH
	cpu.sp.Inc()
	cpu.memory.StackPush(cpu.sp.Read(), pc[0]) // push PCL
	cpu.sp.Inc()
	cpu.memory.StackPush(cpu.sp.Read(), cpu.p.Read()) // push P
	cpu.sp.Inc()

	pcl := cpu.memory.Read([2]byte{0xfe, 0xff})
	pch := cpu.memory.Read([2]byte{0xff, 0xff})

	cpu.pc.Load([2]byte{pcl, pch})

	return brkCost
}

func (cpu *CPU) rti() int {
	cpu.pc.Read()

	cpu.p.Load(cpu.memory.StackPull(cpu.sp.Read())) // pull P
	cpu.sp.Dec()
	pcl := cpu.memory.StackPull(cpu.sp.Read()) // pull PCL
	cpu.sp.Dec()
	pch := cpu.memory.StackPull(cpu.sp.Read()) // pull PCH
	cpu.sp.Dec()

	cpu.pc.Load([2]byte{pcl, pch})

	return rtiCost
}
