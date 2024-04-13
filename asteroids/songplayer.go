package asteroids

import (
	"bytes"
	_ "image/png"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	raudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
)

const (
	sampleRate = 48000
)

func NewAudioContext() *audio.Context {
	return audio.NewContext(sampleRate)
}

// SongPlayer represents the current audio state.
type SongPlayer struct {
	audioContext *audio.Context
	seBytes      []byte
	seCh         chan []byte
	playSong     bool
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

func (p *SongPlayer) Play() {
	p.playSong = true
}

func (p *SongPlayer) Update() error {
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

	// if inpututil.IsKeyJustPressed(ebiten.KeyP) {

	if p.playSong {
		p.playSong = false
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
