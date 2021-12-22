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
	instructionTable map[Instruction]instructionFunc
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
	cpu.p.Reset()

	cpu.Reset()

	for {
		pc := cpu.memory.Read(cpu.pc.Read())
		cpu.ir.Load(pc)

		cpu.dump(false, true, true, true)
		cpu.Step()
	}
}

func (cpu *CPU) Step() {
	cycles := cpu.instructionTable[Instruction(cpu.ir.Read())]()
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

func (cpu *CPU) createInstuctionsTable() map[Instruction]instructionFunc {
	instructionTable := make(map[Instruction]instructionFunc, 256)

	instructionTable[BRK] = cpu.brk
	instructionTable[RTI] = cpu.rti

	instructionTable[JMP_ABS] = cpu.jmpAbs

	instructionTable[STY_ZP] = cpu.styZp
	instructionTable[STA_ZP] = cpu.staZp
	instructionTable[STX_ZP] = cpu.stxZp
	instructionTable[STY_ABS] = cpu.styAbs
	instructionTable[STA_ABS] = cpu.staAbs
	instructionTable[STX_ABS] = cpu.stxAbs
	instructionTable[STA_ZP_Y] = cpu.staZpy
	instructionTable[STY_ZP_X] = cpu.styZpx
	instructionTable[STA_ZP_X] = cpu.staZpx
	instructionTable[STX_ZP_Y] = cpu.stxZpy
	instructionTable[STA_ABS_Y] = cpu.staAbsy
	instructionTable[STA_ABS_X] = cpu.staAbsx

	instructionTable[LDY_IMM] = cpu.ldyImm
	instructionTable[LDX_IMM] = cpu.ldxImm
	instructionTable[LDY_ZP] = cpu.ldyZp
	instructionTable[LDA_ZP] = cpu.ldaZp
	instructionTable[LDX_ZP] = cpu.ldxZp
	instructionTable[LDA_IMM] = cpu.ldaImm
	instructionTable[LDY_ABS] = cpu.ldyAbs
	instructionTable[LDA_ABS] = cpu.ldaAbs
	instructionTable[LDX_ABS] = cpu.ldxAbs

	return instructionTable
}
