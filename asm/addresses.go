package asm

type AddressType int

const (
	ADDR_NORMAL AddressType = iota
	ADDR_LABEL_POSITION
	ADDR_LABEL_DATA
)

func (t AddressType) IsLabel() bool {
	return t == ADDR_LABEL_DATA || t == ADDR_LABEL_POSITION
}

// An address reference.
type AddressRef struct {
	Type AddressType
	Address int
	Value int
}

// Applies the label to an instruction to create a
// labeled instruction.
func (r *AddressRef) Apply(i *Instruction) *Instruction {
	i.Type = INST_LABELED
	i.Label = r
	return i
}

// Creates an address reference, given a concrete
// address.
func CreateAddressRef(address int) *AddressRef {
	return &AddressRef{
		Type: ADDR_NORMAL,
		Address: address,
	}
}

// Creates an address reference for the use of
// position labels.
func CreatePositionLabel() *AddressRef {
	return &AddressRef{
		Type: ADDR_LABEL_POSITION,
		Address: -1,
	}
}

// Creates an address reference for the use of
// data labels.
func CreateDataLabel(value int) *AddressRef {
	return &AddressRef{
		Type: ADDR_LABEL_DATA,
		Address: -1,
		Value: value,
	}
}
