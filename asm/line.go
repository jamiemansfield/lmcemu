package asm

type AddressRef struct {
	address int
}

func (r *AddressRef) Evaluate(originalLabel int) (int, int) {
	var label = originalLabel
	if r.address == -1 {
		label++
		r.address = label
	}
	return r.address, label
}

func CreateAddressRef(a int) *AddressRef {
	return &AddressRef{
		address: a,
	}
}

func CreateLabelRef() *AddressRef {
	return &AddressRef{
		address: -1,
	}
}

type EvaluatedLine struct {
	Opcode Opcode
	Address int
}

func (l *EvaluatedLine) Compile() int {
	return int(l.Opcode) * 100 + l.Address
}

func CreateEvaluatedLine(opcode Opcode, address int) EvaluatedLine {
	return EvaluatedLine{
		Opcode: opcode,
		Address: address,
	}
}

type Line struct {
	Opcode Opcode
	Address *AddressRef
}

func (l *Line) Evaluate(originalLabel int) (EvaluatedLine, int) {
	addr, label := l.Address.Evaluate(originalLabel)
	return EvaluatedLine{
		Opcode: l.Opcode,
		Address: addr,
	}, label
}

func CreateLine(opcode Opcode, address *AddressRef) Line {
	return Line{
		Opcode: opcode,
		Address: address,
	}
}
