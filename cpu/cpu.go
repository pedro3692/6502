package cpu

import (
	"time"

	"github.com/pedro3692/6502/internal/memory"
	"github.com/pedro3692/6502/internal/register"
)

type CPU struct {
	frequency        float32 // in MHz
	memory           memory.Memory
	ir               register.Register
	a                register.Register
	x                register.Register
	y                register.Register
	sp               register.StrackPointer
	pc               register.ProgramCounter
	p                register.ProcessorStatus
	instructionTable map[byte]instructionFunc
	cycleCounter     int64
}

type instructionFunc func() int

func (cpu *CPU) Start(frequency float32) {
	cpu.instructionTable = cpu.createInstuctionsTable()
	cpu.frequency = frequency

	// init registers
	cpu.a = register.New(&cpu.memory)
	cpu.x = register.New(&cpu.memory)
	cpu.y = register.New(&cpu.memory)

	cpu.sp.Reset()
	cpu.pc.Reset()
	cpu.p.Reset()

	for {
		cpu.Step()
		cpu.dump(false, true, true, true)
	}
}

func (cpu *CPU) Step() {
	pc := cpu.memory.Read(cpu.pc.Read())
	cpu.ir.Load(pc)

	cycles := cpu.instructionTable[cpu.ir.Read()]()
	cpu.run(cycles)
}

func (cpu *CPU) Reset() {
	cpu.pc.Reset()

	// load reset vector
	lb := cpu.memory.Read(cpu.pc.Read())
	hb := cpu.memory.Read(cpu.pc.Read())

	cpu.pc.Load([2]byte{lb, hb})
}

func (cpu *CPU) Load(mem [memory.Size]byte) {
	cpu.memory.Set(mem)
}

func (cpu *CPU) run(cycles int) {
	time.Sleep(time.Duration(float32(cycles)*cpu.frequency) * time.Microsecond)
	cpu.cycleCounter += int64(cycles)
}

func (cpu *CPU) createInstuctionsTable() map[byte]instructionFunc {
	instructionTable := make(map[byte]instructionFunc, 256)

	instructionTable[0x00] = cpu.brk

	instructionTable[0x4c] = cpu.jmpAbs

	instructionTable[0x84] = cpu.styZp
	instructionTable[0x85] = cpu.staZp
	instructionTable[0x86] = cpu.stxZp
	instructionTable[0x8c] = cpu.styAbs
	instructionTable[0x8d] = cpu.staAbs
	instructionTable[0x8e] = cpu.stxAbs
	instructionTable[0x91] = cpu.staZpy
	instructionTable[0x94] = cpu.styZpx
	instructionTable[0x95] = cpu.staZpx
	instructionTable[0x96] = cpu.stxZpy
	instructionTable[0x99] = cpu.staAbsy
	instructionTable[0x9D] = cpu.staAbsx

	instructionTable[0xa0] = cpu.ldyImm
	instructionTable[0xa2] = cpu.ldxImm
	instructionTable[0xa4] = cpu.ldyZp
	instructionTable[0xa5] = cpu.ldaZp
	instructionTable[0xa6] = cpu.ldxZp
	instructionTable[0xa9] = cpu.ldaImm
	instructionTable[0xac] = cpu.ldyAbs
	instructionTable[0xad] = cpu.ldaAbs
	instructionTable[0xae] = cpu.ldxAbs

	return instructionTable
}
