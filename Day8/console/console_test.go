package console

import (
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsole_LoadInstruction(t *testing.T) {
	c := NewConsole()
	assert.Equal(t, len(c.Instructions), 0)
	i := NewInstruction("some", 1)
	c.LoadInstruction(i)
	assert.Equal(t, len(c.Instructions), 1)
}

func TestConsole_GetInstruction(t *testing.T) {
	c := NewConsole()

	i := NewInstruction("Test", 1)
	c.LoadInstruction(i)
	i2 := NewInstruction("Test", 2)
	c.LoadInstruction(i2)

	res := c.GetInstruction()

	assert.Equal(t, i, res)
}

func TestConsole_ProcessNop(t *testing.T) {
	c := NewConsole()
	i := NewInstruction("nop", 15)

	c.Process(i)

	assert.Equal(t, c.Index, 1)
	assert.Equal(t, c.Accumulator, 0)
}

func TestConsole_ProcessAcc(t *testing.T) {
	c := NewConsole()
	i := NewInstruction("acc", 15)

	c.Process(i)

	assert.Equal(t, c.Index, 1)
	assert.Equal(t, c.Accumulator, 15)
}

func TestConsole_ProcessJmp(t *testing.T) {
	c := NewConsole()
	i := NewInstruction("jmp", 4)

	c.Process(i)

	assert.Equal(t, c.Index, 4)
	assert.Equal(t, c.Accumulator, 0)
}

func TestConsole_Write(t *testing.T) {
	c := NewConsole()
	err := filehandler.LoadInputFile("testdata/input.txt", c)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	assert.Equal(t, len(c.Instructions), 9)
}

func TestConsole_ProcessMiniRun(t *testing.T) {
	c := NewConsole()
	err := filehandler.LoadInputFile("testdata/input.txt", c)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	i := c.GetInstruction()
	c.Process(i)
	assert.Equal(t, c.Index, 1)
	assert.Equal(t, c.Accumulator, 0)

	i = c.GetInstruction()
	c.Process(i)
	assert.Equal(t, c.Index, 2)
	assert.Equal(t, c.Accumulator, 1)

	i = c.GetInstruction()
	c.Process(i)
	assert.Equal(t, c.Index, 6)
	assert.Equal(t, c.Accumulator, 1)
}

func TestConsole_Run(t *testing.T) {
	c := NewConsole()
	err := filehandler.LoadInputFile("testdata/input.txt", c)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	res := c.Run()

	assert.Equal(t, res, 5)
}

func TestConsole_Revert(t *testing.T) {
	c := NewConsole()
	i := NewInstruction("acc", 15)

	c.Process(i)
	assert.Equal(t, c.Accumulator, 15)
	assert.Equal(t, c.Index, 1)
	assert.Equal(t, i.Reverted, false)

	c.Revert(i)
	assert.Equal(t, c.Accumulator, 0)
	assert.Equal(t, c.Index, 0)
	assert.Equal(t, i.Reverted, true)
}

func TestFiLoQueue_Push(t *testing.T) {
	q := NewFiLoQueue()
	i1 := NewInstruction("Insert1", 1)
	i2 := NewInstruction("Insert2", 2)

	q.Push(i1)
	q.Push(i2)

	assert.Equal(t, q.Items[0], i2)
	assert.Equal(t, q.Items[1], i1)
}

func TestFiLoQueue_Pop(t *testing.T) {
	q := NewFiLoQueue()
	i1 := NewInstruction("Insert1", 1)
	i2 := NewInstruction("Insert2", 2)

	q.Push(i1)
	q.Push(i2)

	res := q.Pop()
	assert.Equal(t, res, i2)
	assert.Equal(t, len(q.Items), 1)
}

func TestInstruction_SwapNop(t *testing.T) {
	i := NewInstruction("nop", 11)

	res := i.Swap()
	assert.Equal(t, i.Operation, "jmp")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, true)
	assert.Equal(t, res, true)
}

func TestInstruction_SwapJmp(t *testing.T) {
	i := NewInstruction("jmp", 11)

	res := i.Swap()
	assert.Equal(t, i.Operation, "nop")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, true)
	assert.Equal(t, res, true)
}

func TestInstruction_SwapAcc(t *testing.T) {
	i := NewInstruction("acc", 11)

	res := i.Swap()
	assert.Equal(t, i.Operation, "acc")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, false)
	assert.Equal(t, res, false)
}

func TestInstruction_SwapTwice(t *testing.T) {
	i := NewInstruction("jmp", 11)

	res := i.Swap()
	assert.Equal(t, i.Operation, "nop")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, true)
	assert.Equal(t, res, true)

	res = i.Swap()
	assert.Equal(t, i.Operation, "nop")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, true)
	assert.Equal(t, res, false)
}

func TestInstruction_RevertNop(t *testing.T) {
	i := NewInstruction("nop", 11)

	res := i.Swap()
	assert.Equal(t, i.Operation, "jmp")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, true)
	assert.Equal(t, res, true)

	i.Revert()
	assert.Equal(t, i.Operation, "nop")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, false)
	assert.Equal(t, i.Reverted, true)
}

func TestInstruction_RevertJmp(t *testing.T) {
	i := NewInstruction("jmp", 11)

	i.Swap()
	assert.Equal(t, i.Operation, "nop")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, true)

	i.Revert()
	assert.Equal(t, i.Operation, "jmp")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, false)
	assert.Equal(t, i.Reverted, true)
}

func TestInstruction_RevertAcc(t *testing.T) {
	i := NewInstruction("acc", 11)

	i.Swap()
	assert.Equal(t, i.Operation, "acc")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, false)

	i.Revert()
	assert.Equal(t, i.Operation, "acc")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, false)
	assert.Equal(t, i.Reverted, false)
}

func TestInstruction_Revert(t *testing.T) {
	i := NewInstruction("jmp", 11)

	i.Swap()
	assert.Equal(t, i.Operation, "nop")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, true)

	i.Revert()
	assert.Equal(t, i.Operation, "jmp")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, false)
	assert.Equal(t, i.Reverted, true)

	i.Revert()
	assert.Equal(t, i.Operation, "jmp")
	assert.Equal(t, i.Value, 11)
	assert.Equal(t, i.Swapped, false)
	assert.Equal(t, i.Reverted, true)
}
