package asteroids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	playerSize   = 80
	playerMargin = 4
)

var (
	tileImage = ebiten.NewImage(playerSize, playerSize)
)

func init() {
	tileImage.Fill(color.White)
}

type PlayerPos struct {
	x int
	y int
}

type Player struct {
	currentPos PlayerPos
	nextPos    PlayerPos
}

func NewPlayer(boardSize int) *Player {
	return &Player{
		currentPos: PlayerPos{
			x: 0,
			y: boardSize - 1,
		},
		nextPos: PlayerPos{
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

	if p.nextPos.x+x >= 0 && p.nextPos.x+x < boardSize {
		p.nextPos.x += x
	}

	if p.nextPos.y+y >= 0 && p.nextPos.y+y < boardSize {
		p.nextPos.y += y
	}
}

func (p *Player) Draw(boardImage *ebiten.Image) {
	ni, nj := p.nextPos.x, p.nextPos.y

	op := &ebiten.DrawImageOptions{}

	nx := ni*playerSize + (ni+1)*playerMargin
	ny := nj*playerSize + (nj+1)*playerMargin

	op.GeoM.Translate(float64(nx), float64(ny))
	op.ColorScale.ScaleWithColor(frameColor)
	boardImage.DrawImage(tileImage, op)
}
