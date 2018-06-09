package asm

func AssembleProgram(lines [100]Line) [100]int {
	var compiled = [100]int{}
	for k, line := range lines {
		compiled[k] = line.Compile()
	}
	return compiled
}
