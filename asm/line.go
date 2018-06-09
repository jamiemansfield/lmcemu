package asm

type AddressRef struct {
	Address int
}

func (r *AddressRef) EvalAddress(label *int) int {
	if r.Address == -1 {
		*label++
		r.Address = *label
	}
	return r.Address
}

func CreateAddressRef(a int) *AddressRef {
	return &AddressRef{
		Address: a,
	}
}

func CreateLabelRef() *AddressRef {
	return &AddressRef{
		Address: -1,
	}
}

type Line struct {
	Opcode Opcode
	Address *AddressRef
}

func (l *Line) Compile(label *int) int {
	return int(l.Opcode) * 100 + l.Address.EvalAddress(label)
}

func CreateLine(opcode Opcode, address *AddressRef) Line {
	return Line{
		Opcode: opcode,
		Address: address,
	}
}

func HLT() Line {
	return CreateLine(OP_HLT, CreateAddressRef(0))
}

func ADD(address *AddressRef) Line {
	return CreateLine(OP_ADD, address)
}

func SUB(address *AddressRef) Line {
	return CreateLine(OP_SUB, address)
}

func STA(address *AddressRef) Line {
	return CreateLine(OP_STA, address)
}

func INP() Line {
	return CreateLine(OP_INP_OUT, CreateAddressRef(1))
}

func OUT() Line {
	return CreateLine(OP_INP_OUT, CreateAddressRef(2))
}
