package new_asm

func DAT(address *AddressRef) *Instruction {
	return CreateInstruction(OP_DAT, address)
}

func HLT() *Instruction {
	return CreateInstruction(OP_HLT, CreateAddressRef(0))
}

func ADD(address *AddressRef) *Instruction {
	return CreateInstruction(OP_ADD, address)
}

func SUB(address *AddressRef) *Instruction {
	return CreateInstruction(OP_SUB, address)
}

func STA(address *AddressRef) *Instruction {
	return CreateInstruction(OP_STA, address)
}

func LDA(address *AddressRef) *Instruction {
	return CreateInstruction(OP_LDA, address)
}

func BRA(address *AddressRef) *Instruction {
	return CreateInstruction(OP_BRA, address)
}

func BRZ(address *AddressRef) *Instruction {
	return CreateInstruction(OP_BRZ, address)
}

func BRP(address *AddressRef) *Instruction {
	return CreateInstruction(OP_BRP, address)
}

func INP() *Instruction {
	return CreateInstruction(OP_INP_OUT, CreateAddressRef(1))
}

func OUT() *Instruction {
	return CreateInstruction(OP_INP_OUT, CreateAddressRef(2))
}

