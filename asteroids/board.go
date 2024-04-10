package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Board represents the game board.
type Board struct {
	size   int
	player *Player
	shots  []*Shot
	// rocks
}

// NewBoard generates a new Board with giving a size.
func NewBoard(size int) *Board {
	b := &Board{
		size:   size,
		player: NewPlayer(size),
		shots:  make([]*Shot, 0),
		// rocks
	}

	// Add random rocks

	return b
}

// Update updates the board state.
func (b *Board) Update(input *Input) error {
	// Update rocks

	shotsToRemove := make([]int, 0)
	for idx, s := range b.shots {
		if alive := s.Move(boardSize); !alive {
			shotsToRemove = append(shotsToRemove, idx)
		}
	}

	for _, idx := range shotsToRemove {
		b.RemoveShot(idx)
	}

	if dir, ok := input.Dir(); ok {
		b.Move(dir)
	}

	if action, ok := input.Action(); ok {
		b.TakeAction(action)
	}

	return nil
}

func (b *Board) Move(dir Direction) {
	b.player.Move(dir, b.size)
}

func (b *Board) TakeAction(action Action) {
	switch action {
	case ActionShot:
		b.AddShot()
	}
}

func (b *Board) AddShot() {
	x, y := b.player.Pos()
	b.shots = append(b.shots, NewShot(x, y))
}

func (b *Board) RemoveShot(idx int) {
	// Se o tamanho é maior que idx, significa que idx existe
	if len(b.shots) > idx {
		b.shots[idx] = b.shots[len(b.shots)-1]
		b.shots = b.shots[:len(b.shots)-1]
	}
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
	for _, s := range b.shots {
		s.Draw(boardImage)
	}

	b.player.Draw(boardImage)
}
