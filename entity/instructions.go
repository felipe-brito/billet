package entity

// Define a instructions type
// @instructions is private variable with the instructions of billet
type Instructions struct {
	instructions []string
}

// Put a new instruction on the list of instructions
func (i *Instructions) add(instruction string) {
	if len(i.instructions) <= 0 {
		i.instructions = make([]string, 6)
	}
	if len(i.instructions) < 6 {
		i.instructions = append(i.instructions, instruction)
	}
}
