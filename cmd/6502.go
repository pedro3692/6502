package main

import (
	"github.com/pedro3692/6502/cpu"
	"github.com/pedro3692/6502/internal/memory"
)

func main() {
	var cpu cpu.CPU

	cpu.Frequency = 1.66
	//test mem
	var mem [memory.Size]byte
	mem[0x0000] = 0x9a
	mem[0x0001] = 0x45
	mem[0x0002] = 0x8d
	mem[0x0002] = 0x55
	mem[0x0002] = 0x56
	mem[0xFFFC] = 0x4c
	mem[0xFFFD] = 0x00
	mem[0xFFFE] = 0x00
	cpu.Load(mem)

	cpu.Start()
}
