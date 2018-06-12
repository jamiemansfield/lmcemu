package asm

func AssembleProgram(lines []*Instruction) ([100]int, error) {
	var compiled = [100]int{}

	// Evaluate DAT operations
	for k, line := range lines {
		if line.Opcode == OP_DAT {
			line.AddressRef.Address = k
		} else
		if line.Type == INST_LABELED {
			line.Label.Address = k
		}
	}

	// Compile
	for k, line := range lines {
		inst, err := line.Evaluate()
		if err != nil {
			return [100]int{}, err
		}
		compiled[k] = inst.Compile()
	}

	return compiled, nil
}
