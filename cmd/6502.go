package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pedro3692/6502/cpu"
	"github.com/pedro3692/6502/internal/memory"
)

func main() {
	var (
		cpu cpu.CPU
		mem [memory.Size]byte
	)

	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("error reading input file %s", err.Error())
	}

	if len(input) != memory.Size {
		log.Fatalln("error invalid file")
	}

	copy(mem[:], input[:])

	// mem[0x0000] = 0xa9
	// mem[0x0001] = 0x45
	// mem[0x0002] = 0x8d
	// mem[0x0003] = 0x55
	// mem[0x0004] = 0x56
	// mem[0xFFFC] = 0x4c
	// mem[0xFFFD] = 0x00
	// mem[0xFFFE] = 0x00
	// mem[0xFFFF] = 0x00

	fmt.Printf("[0x0000:0x000F] = %x\n", mem[0x0000:0x000F+1])
	fmt.Printf("[0x8000:0x800F] = %x\n", mem[0x8000:0x800F+1])
	fmt.Printf("[0x8F00:0x8F0F] = %x\n", mem[0x8F00:0x8F0F+1])
	fmt.Printf("[0xFFF0:0xFFFF] = %x\n", mem[0xFFF0:0xFFFF+1])
	fmt.Printf("[0xFFFC] = %x\n", mem[0xFFFC])
	fmt.Printf("[0xFFFD] = %x\n", mem[0xFFFD])
	fmt.Printf("[0xFFFE] = %x\n", mem[0xFFFE])
	fmt.Printf("[0xFFFF] = %x\n", mem[0xFFFF])

	cpu.Load(mem)
	cpu.Start(1.66)
}
