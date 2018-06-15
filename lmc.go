package main

import (
	"fmt"
	"github.com/jamiemansfield/lmcemu/asm"
	"log"
	"github.com/jamiemansfield/lmcemu/emu"
)

func main() {
	// Assemble the program.
	prog, err := asm.AssembleProgram(Calculator)
	if err != nil {
		log.Fatal(err)
		return
	}

	debug := false

	// Debug assembly
	if debug {
		fmt.Printf("%v", prog)
	}

	// Run
	if !debug {
		cpu := emu.CreateLmcCpu()
		memory := emu.CreateMemory(prog)
		if cpu.Execute(memory) != nil {
			log.Fatal(err)
			return
		}
	}
}
