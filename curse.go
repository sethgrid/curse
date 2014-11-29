package curse

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Cursor struct {
	Position
	Color
}

type Position struct {
	X, Y int
}

type Color struct{}

func (c *Cursor) MoveUp(nLines int) *Cursor {
	fmt.Printf("%c[%dA", ESC, nLines)
	c.Position.Y -= nLines
	return c
}

func (c *Cursor) MoveDown(nLines int) *Cursor {
	fmt.Printf("%c[%dB", ESC, nLines)
	c.Position.Y += nLines
	return c
}

func (c *Cursor) EraseCurrentLine() *Cursor {
	fmt.Printf("%c[2K\r", ESC)
	c.Position.X = 1
	return c
}

func (c *Cursor) Reset() *Cursor {
	if c.Position.Y < 0 {
		c.MoveDown(-1 * c.Position.Y)
	} else if c.Position.Y == 0 {
		// no-op
	} else {
		c.MoveUp(c.Position.Y)
	}
	c.Position.Y = 0
	c.Position.X = 1
	return c
}

func setRawMode() {
	rawMode := exec.Command("/bin/stty", "raw")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()
}

func setCookedMode() {
	cookedMode := exec.Command("/bin/stty", "-raw")
	cookedMode.Stdin = os.Stdin
	_ = cookedMode.Run()
	cookedMode.Wait()
}

func GetCursorPosition() (int, int, error) {
	// set terminal to raw mode
	setRawMode()

	// same as $ echo -e "\033[6n"
	// by printing the output, we are triggering input
	fmt.Printf(fmt.Sprintf("%c[6n", 27))

	// capture keyboard output from print command
	reader := bufio.NewReader(os.Stdin)

	// capture the triggered stdin from the print
	text, _ := reader.ReadSlice('R')

	// restore the terminal mode
	setCookedMode()

	// check for the desired output
	re := regexp.MustCompile(`\d+;\d+`)
	res := re.FindString(string(text))
	if res != "" {
		parts := strings.Split(res, ";")
		line, _ := strconv.Atoi(parts[0])
		col, _ := strconv.Atoi(parts[1])
		return line, col, nil

	} else {
		return 0, 0, errors.New("unable to read cursor position")
	}
}

const (
	ESC = 27
)
