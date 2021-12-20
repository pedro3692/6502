package cpu

import (
	"time"

	"github.com/pedro3692/6502/internal/memory"
	"github.com/pedro3692/6502/internal/register"
)

type CPU struct {
	frequency float32 // in MHz
	memory    memory.Memory
	ir        register.Register
	a         register.Register
	x         register.Register
	y         register.Register
	sp        register.StrackPointer
	pc        register.ProgramCounter
	p         register.ProcessorStatus
}

type instructionFunc func() int

func (cpu *CPU) Start(frequency float32) {
	cpu.frequency = frequency
	cpu.Reset()

	instructionTable := cpu.createInstuctionsTable()

	for {
		pc := cpu.memory.Read(cpu.pc.Read())
		cpu.ir.Load(pc)

		cycles := instructionTable[cpu.ir.Read()]()
		cpu.run(cycles)
	}
}

func (cpu *CPU) Reset() {
	// init registers
	cpu.a = register.New(&cpu.memory)
	cpu.x = register.New(&cpu.memory)
	cpu.y = register.New(&cpu.memory)

	cpu.sp.Reset()
	cpu.pc.Reset()

	lb := cpu.memory.Read(cpu.pc.Read())
	hb := cpu.memory.Read(cpu.pc.Read())
	cpu.pc.Load([2]byte{lb, hb})
}

func (cpu *CPU) Load(mem [memory.Size]byte) {
	cpu.memory.Set(mem)
}

func (cpu CPU) run(cycles int) {
	time.Sleep(time.Duration(float32(cycles)*cpu.frequency) * time.Microsecond)
}

func (cpu *CPU) createInstuctionsTable() map[byte]instructionFunc {
	instructionTable := make(map[byte]instructionFunc, 256)

	instructionTable[0x00] = cpu.brk

	instructionTable[0x4c] = cpu.jmpAbs

	instructionTable[0x8c] = cpu.styAbs
	instructionTable[0x8d] = cpu.staAbs
	instructionTable[0x8e] = cpu.stxAbs

	instructionTable[0xa0] = cpu.ldyIm
	instructionTable[0xa2] = cpu.ldxIm
	instructionTable[0xa9] = cpu.ldaIm

	return instructionTable
}
