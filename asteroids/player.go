package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	playerSize   = 80
	playerMargin = 4
)

var playerImage *ebiten.Image

func init() {
	playerImage = spriteSheet.Crown
}

type PlayerPos struct {
	x int
	y int
}

type Player struct {
	currentPos PlayerPos
}

func NewPlayer(boardSize int) *Player {
	return &Player{
		currentPos: PlayerPos{
			x: 0,
			y: boardSize - 1,
		},
	}
}

func (p *Player) Pos() (int, int) {
	return p.currentPos.x, p.currentPos.y
}

func (p *Player) Move(dir Direction, boardSize int) {
	x, y := dir.Vector()

	if p.currentPos.x+x >= 0 && p.currentPos.x+x < boardSize {
		p.currentPos.x += x
	}

	if p.currentPos.y+y >= 0 && p.currentPos.y+y < boardSize {
		p.currentPos.y += y
	}
}

func (p *Player) Draw(boardImage *ebiten.Image) {
	ni, nj := p.currentPos.x, p.currentPos.y

	op := &ebiten.DrawImageOptions{}

	nx := ni*playerSize + (ni+1)*playerMargin
	ny := nj*playerSize + (nj+1)*playerMargin

	op.GeoM.Translate(float64(nx), float64(ny))
	boardImage.DrawImage(playerImage, op)
}
