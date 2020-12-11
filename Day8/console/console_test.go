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