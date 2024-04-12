// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package asteroids

import (
	"bytes"
	"image"
	_ "image/png"

	"asteroids-go/resources/images"

	"github.com/hajimehoshi/ebiten/v2"
)

const tileSize = 64

// SpriteSheet represents a collection of sprite images.
type SpriteSheet struct {
	Crown *ebiten.Image
	Rocks []*ebiten.Image
	Shot  *ebiten.Image
}

// LoadSpriteSheet loads the embedded SpriteSheet.
func LoadSpriteSheet() (*SpriteSheet, error) {
	img, _, err := image.Decode(bytes.NewReader(images.Spritesheet_png))
	if err != nil {
		return nil, err
	}

	sheet := ebiten.NewImageFromImage(img)

	// spriteAt returns a sprite at the provided coordinates.
	spriteAt := func(x, y int) *ebiten.Image {
		return sheet.SubImage(image.Rect(x*tileSize, (y+1)*tileSize, (x+1)*tileSize, y*tileSize)).(*ebiten.Image)
	}

	// Populate SpriteSheet.
	s := &SpriteSheet{}
	s.Crown = spriteAt(8, 6)

	s.Rocks = append(s.Rocks, spriteAt(3, 2))
	s.Rocks = append(s.Rocks, spriteAt(5, 2))
	s.Rocks = append(s.Rocks, spriteAt(7, 2))
	s.Rocks = append(s.Rocks, spriteAt(9, 2))

	s.Shot = spriteAt(7, 4)

	return s, nil
}
