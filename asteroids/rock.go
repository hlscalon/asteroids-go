package asteroids

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	rockSize   = 80
	rockMargin = 4
)

var rockExploding *ebiten.Image

func init() {
	rockExploding = spriteSheet.RockExplosion
}

type RockPos struct {
	x int
	y int
}

type Rock struct {
	currentPos     RockPos
	isAlive        bool
	image          *ebiten.Image
	isExploding    bool
	explodingCount int
	musicPlayer    *SongPlayer
}

func NewRock(playerX, playerY int, audioContext *audio.Context) *Rock {
	x := playerX
	y := playerY

	musicPlayer, err := NewSongPlayer(audioContext)
	if err != nil {
		log.Fatalf("Error song player: %v\n", err)
		return nil
	}

	return &Rock{
		currentPos:     RockPos{x, y},
		isAlive:        true,
		image:          spriteSheet.Rocks[rand.Intn(len(spriteSheet.Rocks))],
		isExploding:    false,
		explodingCount: 5,
		musicPlayer:    musicPlayer,
	}
}

func (r *Rock) IsAlive() bool {
	return r.isAlive
}

func (r *Rock) SetIsAlive(isAlive bool) {
	r.isAlive = isAlive
}

func (r *Rock) IsExploding() bool {
	return r.isExploding
}

func (r *Rock) SetIsExploding(isExploding bool) {
	r.isExploding = isExploding
	r.musicPlayer.PlayOnUpdate()
}

func (r *Rock) ExplodingCount() int {
	r.explodingCount--

	return r.explodingCount
}

func (r *Rock) Pos() (int, int) {
	return r.currentPos.x, r.currentPos.y
}

func (r *Rock) Move(boardSize int) bool {
	delta := rand.Intn(20)
	if delta != 1 {
		return true
	}

	if r.currentPos.y+delta < boardSize {
		r.currentPos.y += delta

		return true
	}

	return false
}

func (r *Rock) Update() {
	if r.musicPlayer != nil {
		if err := r.musicPlayer.Update(); err != nil {
			log.Printf("Error updating music player: %v\n", err)
		}
	}
}

func (r *Rock) Draw(boardImage *ebiten.Image) {
	ni, nj := r.currentPos.x, r.currentPos.y
	nx := ni*rockSize + (ni+1)*rockMargin
	ny := nj*rockSize + (nj+1)*rockMargin

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(nx), float64(ny))

	if r.isExploding {
		boardImage.DrawImage(rockExploding, op)
	} else {
		boardImage.DrawImage(r.image, op)
	}
}
