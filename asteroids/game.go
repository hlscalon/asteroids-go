package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 680
	boardSize    = 8
)

// Game represents a game state.
type Game struct {
	input       *Input
	board       *Board
	boardImage  *ebiten.Image
	musicPlayer *SongPlayer
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	m, err := NewSongPlayer(audio.NewContext(sampleRate))
	if err != nil {
		return nil, err
	}

	g := &Game{
		input:       NewInput(),
		musicPlayer: m,
	}

	g.board = NewBoard(boardSize)

	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	if err := g.board.Update(g.input); err != nil {
		return err
	}

	if g.musicPlayer != nil {
		if err := g.musicPlayer.update(); err != nil {
			return err
		}
	}

	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}

	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
