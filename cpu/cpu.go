package cpu

import (
	"time"

	"github.com/pedro3692/6502/internal/register"
)

const BusSize uint = 0xFFFF + 0x1

type Bus interface {
	Store([2]byte, byte)
	Read([2]byte) byte
	StackPush(byte, byte)
	StackPull(byte) byte
	Dump(uint16, uint16) []byte
}

type CPU struct {
	frequency        float32 // in MHz
	bus              Bus
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
	cpu.a = register.New(cpu.bus)
	cpu.x = register.New(cpu.bus)
	cpu.y = register.New(cpu.bus)

	cpu.sp.Reset()
	cpu.p.Reset()

	cpu.Reset()

	for {
		pc := cpu.bus.Read(cpu.pc.Read())
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
	lb := cpu.bus.Read(cpu.pc.Read())
	hb := cpu.bus.Read(cpu.pc.Read())

	cpu.pc.Load([2]byte{lb, hb})
}

func (cpu *CPU) run(cycles int) {
	time.Sleep(time.Duration(float32(cycles)*cpu.frequency) * time.Microsecond)
	cpu.cycleCounter += int64(cycles)
}

func New(bus Bus) CPU {
	return CPU{
		bus: bus,
	}
}
