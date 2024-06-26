package asteroids

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var spriteSheet *SpriteSheet

func init() {
	var err error
	spriteSheet, err = LoadSpriteSheet()
	if err != nil {
		panic(fmt.Errorf("failed to load embedded spritesheet: %s", err))
	}
}

type Board struct {
	size         int
	player       *Player
	shots        []*Shot
	rocks        []*Rock
	score        *Score
	audioContext *audio.Context

	lastRockInserted time.Time
}

func NewBoard(size int, audioContext *audio.Context) *Board {
	return &Board{
		size:             size,
		player:           NewPlayer(size),
		score:            NewScore(),
		shots:            make([]*Shot, 0, 20), // pré aloca 20 tiros
		rocks:            make([]*Rock, 0, 20), // pŕe aloca 20 rochas
		lastRockInserted: time.Now(),
		audioContext:     audioContext,
	}
}

func (b *Board) RemoveShots() {
	shotsToRemove := make([]int, 0)
	for idx, s := range b.shots {
		if !s.IsAlive() {
			shotsToRemove = append(shotsToRemove, idx)
		}
	}

	for _, idx := range shotsToRemove {
		b.RemoveShot(idx)
	}
}

func (b *Board) MoveShots() {
	for _, s := range b.shots {
		if s.IsAlive() {
			s.SetIsAlive(s.Move())
		}
	}
}

func (b *Board) RemoveRocks() {
	rocksToRemove := make([]int, 0)
	for idx, r := range b.rocks {
		if r.IsExploding() {
			if r.ExplodingCount() == 0 {
				r.SetIsExploding(false)
			}
		} else {
			if !r.IsAlive() {
				rocksToRemove = append(rocksToRemove, idx)
			}
		}
	}

	for _, idx := range rocksToRemove {
		b.RemoveRock(idx)
	}
}

func (b *Board) MoveRocks() {
	for _, r := range b.rocks {
		if r.IsAlive() {
			r.SetIsAlive(r.Move(boardSize))
		}
	}
}

func (b *Board) DetectCollisions() {
	ch := make(chan int)

	go func() {
		for _, s := range b.shots {
			for _, r := range b.rocks {
				shotX, shotY := s.Pos()
				rockX, rockY := r.Pos()

				if shotX == rockX && shotY == rockY {
					s.SetIsAlive(false)
					r.SetIsAlive(false)
					r.SetIsExploding(true)

					log.Printf("Colisão shot|rock: %d, %d", shotX, shotY)

					b.score.Add()
				}
			}
		}

		ch <- 1
	}()

	go func() {
		playerX, playerY := b.player.Pos()
		for _, r := range b.rocks {
			rockX, rockY := r.Pos()

			if rockX == playerX && rockY == playerY {
				log.Printf("Jogador morreu: %d, %d", rockX, rockY)
				b.score.Reset()
				r.SetIsAlive(false)

				break
			}
		}

		ch <- 1
	}()

	<-ch
	<-ch
}

func (b *Board) MoveEntities(input *Input) {
	ch := make(chan int)

	go func() {
		b.MoveShots()

		ch <- 1
	}()

	go func() {
		b.MoveRocks()

		ch <- 1
	}()

	go func() {
		if dir, ok := input.Dir(); ok {
			b.Move(dir)
		}

		ch <- 1
	}()

	<-ch
	<-ch
	<-ch
}

func (b *Board) RemoveEntities() {
	ch := make(chan int)

	go func() {
		b.RemoveShots()

		ch <- 1
	}()

	go func() {
		b.RemoveRocks()
		ch <- 1
	}()

	<-ch
	<-ch
}

func (b *Board) UpdateEntities() {
	for _, r := range b.rocks {
		r.Update()
	}
}

func (b *Board) Update(input *Input) error {
	b.DetectCollisions()
	b.UpdateEntities()
	b.MoveEntities(input)
	b.RemoveEntities()
	b.AddRandomRock()

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

	b.rocks = append(b.rocks, NewRock(x, y, b.audioContext))

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

	b.score.Draw(boardImage)
	b.player.Draw(boardImage)
}
