package cpu

const (
	brkCost = 7
	rtiCost = 6
	nmiCost = 7
	irqCost = 7
)

func (cpu *CPU) brk() int {
	cpu.p.SetBreak()

	cpu.pc.Read()
	pc := cpu.pc.Read() // PC + 2

	cpu.bus.StackPush(cpu.sp.Read(), pc[1]) // push PCH
	cpu.sp.Inc()
	cpu.bus.StackPush(cpu.sp.Read(), pc[0]) // push PCL
	cpu.sp.Inc()
	cpu.bus.StackPush(cpu.sp.Read(), cpu.p.Read()) // push P
	cpu.sp.Inc()

	cpu.p.ResetDecimalMode()
	cpu.p.SetIRQBDisable()

	pcl := cpu.bus.Read([2]byte{0xfe, 0xff})
	pch := cpu.bus.Read([2]byte{0xff, 0xff})

	cpu.pc.Load([2]byte{pcl, pch})

	return brkCost
}

func (cpu *CPU) rti() int {
	cpu.pc.Read()

	cpu.sp.Dec()
	cpu.p.Load(cpu.bus.StackPull(cpu.sp.Read())) // pull P
	cpu.sp.Dec()
	pcl := cpu.bus.StackPull(cpu.sp.Read()) // pull PCL
	cpu.sp.Dec()
	pch := cpu.bus.StackPull(cpu.sp.Read()) // pull PCH

	cpu.p.ResetBreak()

	cpu.pc.Load([2]byte{pcl, pch})

	return rtiCost
}

func (cpu *CPU) nmi() int {
	cpu.p.ResetBreak()

	pc := cpu.pc.Check()

	cpu.bus.StackPush(cpu.sp.Read(), pc[1]) // push PCH
	cpu.sp.Inc()
	cpu.bus.StackPush(cpu.sp.Read(), pc[0]) // push PCL
	cpu.sp.Inc()
	cpu.bus.StackPush(cpu.sp.Read(), cpu.p.Read()) // push P
	cpu.sp.Inc()

	cpu.p.SetIRQBDisable()

	pcl := cpu.bus.Read([2]byte{0xfa, 0xff})
	pch := cpu.bus.Read([2]byte{0xfb, 0xff})

	cpu.pc.Load([2]byte{pcl, pch})

	return nmiCost
}

func (cpu *CPU) irq() int {
	cpu.p.ResetBreak()

	pc := cpu.pc.Check()

	cpu.bus.StackPush(cpu.sp.Read(), pc[1]) // push PCH
	cpu.sp.Inc()
	cpu.bus.StackPush(cpu.sp.Read(), pc[0]) // push PCL
	cpu.sp.Inc()
	cpu.bus.StackPush(cpu.sp.Read(), cpu.p.Read()) // push P
	cpu.sp.Inc()

	cpu.p.SetIRQBDisable()

	pcl := cpu.bus.Read([2]byte{0xfe, 0xff})
	pch := cpu.bus.Read([2]byte{0xff, 0xff})

	cpu.pc.Load([2]byte{pcl, pch})

	return irqCost
}
