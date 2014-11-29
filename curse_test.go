package curse

import (
	"testing"
)

func TestMovementY(t *testing.T) {
	c := &Cursor{}
	c.MoveDown(2)
	c.MoveUp(5)

	if c.Position.Y != -3 {
		t.Errorf("Wrong Y position - got %d, want %d", c.Position.Y, -3)
	}
	c.Reset()
}

func TestChainedMovementY(t *testing.T) {
	c := &Cursor{}
	c.MoveUp(1).MoveUp(1)

	if c.Position.Y != -2 {
		t.Errorf("Wrong Y position - got %d, want %d", c.Position.Y, -2)
	}
	c.Reset()
}

func TestEraseCurrentLine(t *testing.T) {
	c := &Cursor{}
	c.X = 12
	c.EraseCurrentLine()

	if c.Position.X != 1 {
		t.Errorf("Wrong X position - got %d, want %d", c.Position.X, 1)
	}
	c.Reset()
}
