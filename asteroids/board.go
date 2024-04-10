package asteroids

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	size   int
	player *Player
	shots  []*Shot
	rocks  []*Rock

	lastRockInserted time.Time
}

func NewBoard(size int) *Board {
	return &Board{
		size:             size,
		player:           NewPlayer(size),
		shots:            make([]*Shot, 0, 20), // pré aloca 20 tiros
		rocks:            make([]*Rock, 0, 20), // pŕe aloca 20 rochas
		lastRockInserted: time.Now(),
	}
}

func (b *Board) Update(input *Input) error {
	shotsToRemove := make([]int, 0)
	for idx, s := range b.shots {
		if alive := s.Move(); !alive {
			shotsToRemove = append(shotsToRemove, idx)
		}
	}

	for _, idx := range shotsToRemove {
		b.RemoveShot(idx)
	}

	rocksToRemove := make([]int, 0)
	for idx, r := range b.rocks {
		if alive := r.Move(boardSize); !alive {
			rocksToRemove = append(rocksToRemove, idx)
		}
	}

	for _, idx := range rocksToRemove {
		b.RemoveRock(idx)
	}

	b.AddRandomRock()

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

func (b *Board) AddRandomRock() {
	if b.lastRockInserted.Add(time.Duration(1+rand.Intn(3)) * time.Second).After(time.Now()) {
		return
	}

	x := rand.Intn(boardSize)
	y := rand.Intn(boardSize / 3)

	b.rocks = append(b.rocks, NewRock(x, y))

	b.lastRockInserted = time.Now()
}

func (b *Board) RemoveRock(idx int) {
	// Se o tamanho é maior que idx, significa que idx existe
	if len(b.rocks) > idx {
		b.rocks[idx] = b.rocks[len(b.rocks)-1]
		b.rocks = b.rocks[:len(b.rocks)-1]
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

	for _, s := range b.shots {
		s.Draw(boardImage)
	}

	for _, r := range b.rocks {
		r.Draw(boardImage)
	}

	b.player.Draw(boardImage)
}
