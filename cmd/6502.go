package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pedro3692/6502/cpu"
	"github.com/pedro3692/6502/internal/memory"
)

func main() {
	var m [memory.Size]byte

	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("error reading input file %s", err.Error())
	}

	if len(input) != memory.Size {
		log.Fatalln("error invalid file")
	}

	copy(m[:], input[:])

	fmt.Printf("[0x0000:0x000F] = %x\n", m[0x0000:0x000F+1])
	fmt.Printf("[0x8000:0x800F] = %x\n", m[0x8000:0x800F+1])
	fmt.Printf("[0x8F00:0x8F0F] = %x\n", m[0x8F00:0x8F0F+1])
	fmt.Printf("[0xFFF0:0xFFFF] = %x\n", m[0xFFF0:0xFFFF+1])
	fmt.Printf("[0xFFFC] = %x\n", m[0xFFFC])
	fmt.Printf("[0xFFFD] = %x\n", m[0xFFFD])
	fmt.Printf("[0xFFFE] = %x\n", m[0xFFFE])
	fmt.Printf("[0xFFFF] = %x\n", m[0xFFFF])

	mem := memory.New(m)
	c := cpu.New(&mem)

	c.Start(1.66)
}
