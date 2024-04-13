package asteroids

import (
	"bytes"
	_ "image/png"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	raudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480

	sampleRate = 48000
)

// SongPlayer represents the current audio state.
type SongPlayer struct {
	audioContext *audio.Context
	seBytes      []byte
	seCh         chan []byte
}

func NewSongPlayer(audioContext *audio.Context) (*SongPlayer, error) {
	songPlayer := &SongPlayer{
		audioContext: audioContext,
		seCh:         make(chan []byte),
	}

	go func() {
		s, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(raudio.Jab_wav))
		if err != nil {
			log.Fatal(err)
			return
		}

		b, err := io.ReadAll(s)
		if err != nil {
			log.Fatal(err)
			return
		}

		songPlayer.seCh <- b
	}()

	return songPlayer, nil
}

func (p *SongPlayer) update() error {
	select {
	case p.seBytes = <-p.seCh:
		close(p.seCh)
		p.seCh = nil
	default:
	}

	p.playSEIfNeeded()

	return nil
}

func (p *SongPlayer) shouldPlaySE() bool {
	if p.seBytes == nil {
		// Bytes for the SE is not loaded yet.
		return false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return true
	}

	return false
}

func (p *SongPlayer) playSEIfNeeded() {
	if !p.shouldPlaySE() {
		return
	}

	sePlayer := p.audioContext.NewPlayerFromBytes(p.seBytes)
	sePlayer.Play()
}

// type Game struct {
// 	musicPlayer *Player
// }

// func NewGame() (*Game, error) {
// 	audioContext := audio.NewContext(sampleRate)

// 	m, err := NewPlayer(audioContext)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Game{
// 		musicPlayer: m,
// 	}, nil
// }

// func (g *Game) Update() error {
// 	if g.musicPlayer != nil {
// 		if err := g.musicPlayer.update(); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	// empty
// 	// interface method
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return screenWidth, screenHeight
// }

// func main() {
// 	ebiten.SetWindowSize(screenWidth, screenHeight)
// 	ebiten.SetWindowTitle("Audio (Ebitengine Demo)")
// 	g, err := NewGame()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := ebiten.RunGame(g); err != nil {
// 		log.Fatal(err)
// 	}
// }
