package main

import (
	"fmt"
	"github.com/jamiemansfield/lmcemu/asm"
	"log"
	"os"
	"encoding/json"
	"github.com/jamiemansfield/lmcemu/emu"
)

func main() {
	// Read asm from file
	file, err := os.Open("calculator.asm")
	defer file.Close()

	// Create a Parser and parse the file
	parser := asm.CreateParser()
	parser.ReadFromFile(file)

	// Assemble the program.
	instructions, err := parser.Assemble()
	if err != nil {
		log.Fatal(err)
		return
	}

	// JSON debuggery
	if false {
		bytes, err := json.MarshalIndent(instructions, "", "\t")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(string(bytes))
		return
	}

	// Assemble the program into RAM
	prog, err := asm.AssembleProgram(instructions)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Debug assembled machine code
	if false {
		fmt.Printf("%v", prog)
		return
	}

	// Run
	cpu := emu.CreateLmcCpu()
	memory := emu.CreateMemory(prog)
	err = cpu.Execute(memory)
	if err != nil {
		log.Fatal(err)
		return
	}
}
