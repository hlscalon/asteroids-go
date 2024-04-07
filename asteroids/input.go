package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Direction int

const (
	DirUp Direction = iota
	DirRight
	DirDown
	DirLeft
)

// String returns a string representing the direction.
func (d Direction) String() string {
	switch d {
	case DirUp:
		return "Up"
	case DirRight:
		return "Right"
	case DirDown:
		return "Down"
	case DirLeft:
		return "Left"
	}
	panic("not reach")
}

// Vector returns a [-1, 1] value for each axis.
func (d Direction) Vector() (x, y int) {
	switch d {
	case DirUp:
		return 0, -1
	case DirRight:
		return 1, 0
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	}
	panic("not reach")
}

// Input represents the current key states.
type Input struct {
}

// NewInput generates a new Input object.
func NewInput() *Input {
	return &Input{}
}

// Direction returns a currently pressed direDirectionection.
// Direction returns false if no direction key is pressed.
func (i *Input) Dir() (Direction, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		return DirUp, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return DirLeft, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return DirRight, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		return DirDown, true
	}
	return 0, false
}
