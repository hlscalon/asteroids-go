package asteroids

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

var errTaskTerminated = errors.New("asteroids: task terminated")

type task func() error

// Board represents the game board.
type Board struct {
	size   int
	player *Player
	tasks  []task
	// shots
	// rocks
}

// NewBoard generates a new Board with giving a size.
func NewBoard(size int) (*Board, error) {
	b := &Board{
		size:   size,
		player: NewPlayer(size),
		// rocks
		// shots
	}

	// Add random rocks

	return b, nil
}

// Update updates the board state.
func (b *Board) Update(input *Input) error {
	// Update rocks
	// Update shots

	if 0 < len(b.tasks) {
		t := b.tasks[0]
		if err := t(); err == errTaskTerminated {
			b.tasks = b.tasks[1:]
		} else if err != nil {
			return err
		}
		return nil
	}

	if dir, ok := input.Dir(); ok {
		if err := b.Move(dir); err != nil {
			return err
		}
	}

	return nil
}

// Move enqueues tile moving tasks.
func (b *Board) Move(dir Direction) error {
	// Move rocks
	// Move shots

	b.player.Move(dir, b.size)

	return nil
}

// Size returns the board size.
func (b *Board) Size() (int, int) {
	x := b.size*playerSize + (b.size+1)*playerMargin
	y := x
	return x, y
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(backgroundColor)

	// Draw rocks
	// Draw shots

	b.player.Draw(boardImage)
}
