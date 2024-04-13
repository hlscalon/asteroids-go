package asteroids

import (
	"bytes"
	_ "image/png"
	"io"
	"log"

	raudio "asteroids-go/resources/audio"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
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

func (p *SongPlayer) PlayOnUpdate() {
	p.playSong = true
}

func (p *SongPlayer) Update() error {
	select {
	case p.seBytes = <-p.seCh:
		close(p.seCh)
		p.seCh = nil
	default:
	}

	if p.seBytes == nil {
		return nil
	}

	if p.playSong {
		p.playSong = false

		sePlayer := p.audioContext.NewPlayerFromBytes(p.seBytes)
		sePlayer.Play()
	}

	return nil
}
