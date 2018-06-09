package asm

type Line struct {
	Opcode Opcode
	Address int
}

func (l *Line) Compile() int {
	return int(l.Opcode) * 100 + l.Address
}

func CreateLine(opcode Opcode, address int) Line {
	return Line{
		Opcode: opcode,
		Address: address,
	}
}
