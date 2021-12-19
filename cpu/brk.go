package cpu

const brkCost = 7

func (cpu *CPU) brk() int {
	cpu.p.SetBreak()
	cpu.p.ResetDecimalMode()
	cpu.p.SetIRQBDisable()

	sp1 := cpu.sp.Inc()
	sp2 := cpu.sp.Inc()

	cpu.pc.Read()
	cpu.pc.Read()
	pc2 := cpu.pc.Read() // PC + 2

	cpu.memory.StackPush(sp1, pc2[0])
	cpu.memory.StackPush(sp2, pc2[1])

	lb := cpu.memory.Read([2]byte{0xfe, 0xff})
	hb := cpu.memory.Read([2]byte{0xff, 0xff})

	cpu.pc.Load([2]byte{lb, hb})

	return brkCost
}
