package asteroids

import (
	"bytes"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	scoreSize   = 80
	scoreMargin = 4
)

var (
	scoreImage      = ebiten.NewImage(scoreSize, scoreSize)
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
	scoreImage.Fill(color.RGBA{0xee, 0xe4, 0xda, 0xff})
}

type Score struct {
	value int
}

func NewScore() *Score {
	return &Score{}
}

func (s *Score) Add() {
	s.value++
}

func (s *Score) Reset() {
	s.value = 0
}

func (s *Score) Draw(boardImage *ebiten.Image) {
	ni, nj := 0.3, 0.3

	nx := ni*scoreSize + (ni+1)*scoreMargin
	ny := nj*scoreSize + (nj+1)*scoreMargin

	str := strconv.Itoa(s.value)

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(nx, ny)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	text.Draw(boardImage, str, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   48,
	}, textOp)

}
