package asm

func AssembleProgram(lines []Line) [100]int {
	var compiled = [100]int{}
	var label = len(lines)
	for k, line := range lines {
		compiled[k] = line.Compile(&label)
	}
	return compiled
}
