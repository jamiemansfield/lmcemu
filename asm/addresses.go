package asm

// An address reference.
type AddressRef struct {
	Address int
}

// Applies the label to an instruction to create a
// labeled instruction.
func (r *AddressRef) Apply(i *Instruction) *Instruction {
	i.Type = LABELED
	i.Label = r
	return i
}

// Creates an address reference, given a concrete
// address.
func CreateAddressRef(address int) *AddressRef {
	return &AddressRef{
		Address: address,
	}
}

// Creates an address reference for the use of
// labels.
func CreateLabelRef() *AddressRef {
	return &AddressRef{
		Address: -1,
	}
}
