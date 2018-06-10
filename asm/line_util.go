package asm

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

func LDA(address *AddressRef) Line {
	return CreateLine(OP_LDA, address)
}

func BRA(address *AddressRef) Line {
	return CreateLine(OP_BRA, address)
}

func BRZ(address *AddressRef) Line {
	return CreateLine(OP_BRZ, address)
}

func BRP(address *AddressRef) Line {
	return CreateLine(OP_BRP, address)
}

func INP() Line {
	return CreateLine(OP_INP_OUT, CreateAddressRef(1))
}

func OUT() Line {
	return CreateLine(OP_INP_OUT, CreateAddressRef(2))
}
