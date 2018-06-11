package new_asm

func AssembleProgram(lines []*Instruction) [100]int {
	var compiled = [100]int{}

	// Evaluate DAT operations
	for k, line := range lines {
		if line.Opcode == OP_DAT && line.Type != EVALUATED {
			line.AddressRef.Address = k
		}
	}

	// Compile
	for k, line := range lines {
		compiled[k] = line.Evaluate().Compile()
	}

	return compiled
}
