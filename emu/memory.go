package emu

// A representation of RAM within an emulated system.
type Memory struct {
	values [100]int
}

// Gets the integer value of a given memory address.
func (m *Memory) GetValue(address int) int {
	return m.values[address]
}

// Sets the integer value of a given memory address.
func (m *Memory) SetValue(address int, value int) {
	m.values[address] = value
}

// Creates memory with the provided initial values
func CreateMemory(initialValues [100]int) *Memory {
	return &Memory{
		values: initialValues,
	}
}
