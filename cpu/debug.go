package cpu

import "fmt"

func (cpu CPU) dump(full bool) {
	if full {
		fmt.Printf(
			"Memory: %x\n", cpu.memory.Dump(),
		)
	}

	fmt.Printf(
		"Zero Page: %x\nStack: %x\n",
		cpu.memory.DumpZeroPage(),
		cpu.memory.DumpStack(),
	)

	fmt.Printf(
		"PC: %x\nA : %x X: %x Y: %x IR: %x SP: %x\n",
		cpu.pc.Check(),
		cpu.a.Read(),
		cpu.x.Read(),
		cpu.y.Read(),
		cpu.ir.Read(),
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
