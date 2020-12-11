package console

import (
	"strconv"
	"strings"
)

type Console struct {
	Instructions []*Instruction
	Index        int
	Accumulator  int
	SwapPoint	 *Instruction
}

func (c *Console) DFSDebug() int {
	q := NewLiFoQueue()

	for {
		if c.Index >= len(c.Instructions) {
			break
		}
		i := c.GetInstruction()
		b := c.Process(i)
		if b {
			for {
				if c.SwapPoint != nil {
					// Backtrack to swap point
					i = q.Pop()
					c.Revert(i)
					if i == c.SwapPoint {
						i.Revert()
						c.SwapPoint = nil
					}
				} else {
					// Backtrack until we can swap
					i = q.Pop()
					c.Revert(i)
					b = i.Swap()
					if b {
						c.SwapPoint = i
						break
					}
				}
			}
		} else {
			// We only want to push into Q if !b
			q.Push(i)
		}
	}
	return c.Accumulator
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
	}
	i.Visited = true
	switch i.Operation {
	case "nop":
		c.Index++
	case "acc":
		c.Accumulator += i.Value
		c.Index++
	case "jmp":
		c.Index += i.Value
	}
	return false
}

func (c *Console) Revert(i *Instruction) {
	i.Visited = false
	switch i.Operation {
	case "nop":
		c.Index--
	case "acc":
		c.Accumulator -= i.Value
		c.Index--
	case "jmp":
		c.Index -= i.Value
	}

}

func (c *Console) Run() int {
	for {
		i := c.GetInstruction()
		b := c.Process(i)
		if b {
			break
		}
	}
	return c.Accumulator
}

func (c *Console) Write(p []byte) (int, error) {
	rb := len(p)
	rawStr := string(p)
	split := strings.Split(rawStr, "\r\n")
	for _, s := range split {
		if s != "" {
			inS := strings.Split(s, " ")
			inN, err := strconv.Atoi(inS[1])
			if err != nil {
				return 0, err
			}
			i := NewInstruction(inS[0], inN)
			c.LoadInstruction(i)
		}
	}
	return rb, nil
}

type Instruction struct {
	Operation string
	Value     int
	Visited   bool
	Reverted  bool
	Swapped   bool
}

func (i *Instruction) Swap() bool {
	if !i.Reverted {
		if !i.Swapped {
			switch i.Operation {
			case "jmp":
				i.Operation = "nop"
				i.Swapped = true
				return true
			case "nop":
				i.Operation = "jmp"
				i.Swapped = true
				return true
			}
		}
	}
	return false
}

func (i *Instruction) Revert() {
	if i.Swapped {
		switch i.Operation {
		case "jmp":
			i.Operation = "nop"
			i.Swapped = false
			i.Reverted = true
		case "nop":
			i.Operation = "jmp"
			i.Swapped = false
			i.Reverted = true
		}
	}
}

type LiFoQueue struct {
	Items []*Instruction
}

func (q *LiFoQueue) Push(i *Instruction) {
	q.Items = append([]*Instruction{i}, q.Items...)
}

func (q *LiFoQueue) Pop() *Instruction {
	i := q.Items[0]
	if len(q.Items) > 1{
		q.Items = q.Items[1:]
	} else {
		q.Items = make([]*Instruction, 0)
	}
	return i
}

func NewConsole() *Console {
	c := Console{
		Instructions: make([]*Instruction, 0),
		Index:        0,
		Accumulator:  0,
		SwapPoint: 	  nil,
	}
	return &c
}

func NewLiFoQueue() *LiFoQueue {
	q := LiFoQueue{Items: make([]*Instruction, 0)}
	return &q
}

func NewInstruction(op string, val int) *Instruction {
	i := Instruction{
		Operation: op,
		Value:     val,
		Visited:   false,
		Reverted:  false,
	}
	return &i
}
