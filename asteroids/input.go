package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Action int
type Direction int

const (
	DirUp Direction = iota
	DirRight
	DirDown
	DirLeft
)

const (
	ActionShot Action = iota
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
type Input struct{}

// NewInput generates a new Input object.
func NewInput() *Input {
	return &Input{}
}

// Direction returns a currently pressed Direction.
// Direction returns false if no direction key is pressed.
func (i *Input) Dir() (Direction, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		return DirUp, true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		return DirLeft, true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		return DirRight, true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return DirDown, true
	}

	return 0, false
}

func (i *Input) Action() (Action, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return ActionShot, true
	}

	return 0, false
}
