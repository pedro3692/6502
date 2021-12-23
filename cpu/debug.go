package cpu

import "fmt"

func (cpu CPU) dump(fullMem, zpstk, registers, cycles bool) {
	if fullMem {
		fmt.Printf(
			"Memory: %x\n", cpu.bus.Dump(0x0, 0xffff),
		)
	}

	if zpstk {
		fmt.Printf(
			"Zero Page: %x\nStack: %x\n",
			cpu.DumpZeroPage(),
			cpu.DumpStack(),
		)
	}

	if registers {
		pc := cpu.pc.Check()
		fmt.Printf(
			"PCH: %02x PCL: %02x\nIR: %02x [%s]\nA : %02x X: %02x Y: %x\nSP: %02x\n",
			pc[0],
			pc[1],
			cpu.ir.Read(),
			Instruction(cpu.ir.Read()),
			cpu.a.Read(),
			cpu.x.Read(),
			cpu.y.Read(),
			cpu.sp.Read(),
		)

		fmt.Printf(
			"N: %t V: %t B: %t D: %t I: %t Z: %t C: %t\n",
			cpu.p.Negative(),
			cpu.p.Overflow(),
			cpu.p.Break(),
			cpu.p.DecimalMode(),
			cpu.p.IRQBDisable(),
			cpu.p.Zero(),
			cpu.p.Carry(),
		)
	}

	if cycles {
		fmt.Printf("cycle: %d\n", cpu.cycleCounter)
	}

}

func (cpu CPU) DumpZeroPage() []byte {
	return cpu.bus.Dump(0x00, 0xff)
}

func (cpu CPU) DumpStack() []byte {
	return cpu.bus.Dump(0x00, 0xff)
}
