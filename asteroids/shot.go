package asteroids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	shotSize   = 80
	shotMargin = 4
)

var (
	shotImage = ebiten.NewImage(shotSize, shotSize)
)

func init() {
	shotImage.Fill(color.Gray16{0xdddd})
}

type ShotPos struct {
	x int
	y int
}

type Shot struct {
	currentPos ShotPos
}

func NewShot(playerX, playerY int) *Shot {
	x := playerX
	y := playerY

	return &Shot{
		currentPos: ShotPos{x, y},
	}
}

func (s *Shot) Move(boardSize int) bool {
	if s.currentPos.y-1 >= 0 {
		s.currentPos.y--

		return true
	}

	return false
}

func (s *Shot) Draw(boardImage *ebiten.Image) {
	ni, nj := s.currentPos.x, s.currentPos.y

	op := &ebiten.DrawImageOptions{}

	nx := ni*shotSize + (ni+1)*shotMargin
	ny := nj*shotSize + (nj+1)*shotMargin

	op.GeoM.Translate(float64(nx), float64(ny))
	op.ColorScale.ScaleWithColor(frameColor)
	boardImage.DrawImage(shotImage, op)
}