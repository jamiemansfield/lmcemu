package asm

func AssembleProgram(lines []Line) [100]int {
	var compiled = [100]int{}
	var label = len(lines)
	for k, line := range lines {
		eval, newLabel := line.Evaluate(label)
		compiled[k] = eval.Compile()
		label = newLabel
	}
	return compiled
}
