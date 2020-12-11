package console

type Console struct {
	Instructions []*Instruction
	Index int
	Accumulator int
}

func (c *Console) GetInstruction() *Instruction {
	return c.Instructions[c.Index]
}

func (c *Console) LoadInstruction(i *Instruction) {
	c.Instructions = append(c.Instructions, i)
}

func (c *Console) Process(i *Instruction) bool {
	if i.Visited {
		return true
	} else {
		i.Visited = true
	}
	switch i.Operation {
	case "nop":
		c.Index++
	case "acc":
		c.Accumulator += i.Value
	case "jmp":
		c.Index += i.Value
	}
	return false
}

type Instruction struct {
	Operation string
	Value int
	Visited bool
}

func NewConsole() *Console {
	c := Console{
		Instructions: make([]*Instruction, 0),
		Index:        0,
		Accumulator:  0,
	}
	return &c
}

func NewInstruction(op string, val int) *Instruction {
	i := Instruction{
		Operation: op,
		Value:     val,
		Visited:   false,
	}
	return &i
}