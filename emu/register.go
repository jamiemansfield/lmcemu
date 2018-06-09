package emu

// A representation of a register within a CPU.
type Register struct {
	Name string
	value int
}

// Gets the integer value of the register.
func (r *Register) GetValue() int {
	return r.value
}

// Sets the integer value of the register.
func (r *Register) SetValue(value int) {
	r.value = value
}

// Gets the integer value of the register returns it, then increments by 1.
func (r *Register) IncrementAndGet() int {
	var oldVal = r.value
	r.value++
	return oldVal
}

// Increments the integer value of the register by 1.
func (r *Register) Increment() {
	r.value++
}

// Creates a register, using the provided data
func CreateRegister(name string, initialValue int) *Register {
	return &Register{
		Name: name,
		value: initialValue,
	}
}
