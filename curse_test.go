package curse

import "testing"

func TestMovementY(t *testing.T) {
	c := &Cursor{}
	c.MoveDown(2)
	c.MoveUp(5)
	newPosition := c.Position
	c.MoveDown(5).MoveUp(2) // restore

	if newPosition.Y != -3 {
		t.Errorf("Wrong Y position - got %d, want %d", newPosition, -3)
	}
}

func TestChainedMovementY(t *testing.T) {
	c := &Cursor{}
	c.MoveDown(1).MoveDown(1)
	newPosition := c.Position
	c.MoveUp(2) // restore

	if newPosition.Y != 2 {
		t.Errorf("Wrong Y position - got %d, want %d", newPosition.Y, 2)
	}
}

func TestEraseCurrentLine(t *testing.T) {
	c := &Cursor{}
	c.X = 12
	c.EraseCurrentLine()

	if c.Position.X != 1 {
		t.Errorf("Wrong X position - got %d, want %d", c.Position.X, 1)
	}
}
