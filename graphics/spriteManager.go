package graphics

import (
	"fmt"
	"time"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

type SpriteManager struct {
	txtrmgr   TextureManager
	sprites   []Drawable
	animables []Animable
}

func NexSpriteManager(txtrmgr TextureManager) SpriteManager {
	return SpriteManager{
		txtrmgr: txtrmgr,
		sprites: []Drawable{},
	}
}

func (s *SpriteManager) Draw(w sfml.Struct_SS_sfRenderWindow) {
	for i := range s.sprites {
		s.sprites[i].Draw(w)
	}
}

func (s *SpriteManager) Animate(elapsed time.Duration) {
	for i := range s.animables {
		s.animables[i].Animate(elapsed)
	}
}

func (s *SpriteManager) LoadBasicSprite(spriteName, textureName string, textureRect Rect) error {
	txtr, err := s.txtrmgr.GetSfTexture(textureName)
	if err != nil {
		return err
	}
	spr := sfml.SfSprite_create()
	sfml.SfSprite_setTexture(spr, *txtr, 1)

	sfml.SfSprite_setTextureRect(spr, textureRect.ToSFMLIntRect())

	s.sprites = append(s.sprites, BasicSprite{sfml_sprite: &spr, name: spriteName, abspos: Vec2{0, 0}})

	return nil
}

func (s *SpriteManager) LoadAnimatedSprite(spriteName, textureName string, playbackMode PlaybackMode, frames []AnimatedSpriteFrame) error {
	if len(frames) == 0 {
		return fmt.Errorf("Animated sprites must have at least one frame")
	}

	txtr, err := s.txtrmgr.GetSfTexture(textureName)
	if err != nil {
		return err
	}
	spr := sfml.SfSprite_create()
	sfml.SfSprite_setTexture(spr, *txtr, 1)

	sfml.SfSprite_setTextureRect(spr, frames[0].Rect.ToSFMLIntRect())

	sprite := AnimatedSprite{BasicSprite: BasicSprite{sfml_sprite: &spr, name: spriteName, abspos: Vec2{0, 0}}, frames: frames, playbackMode: playbackMode}
	s.sprites = append(s.sprites, &sprite)
	s.animables = append(s.animables, &sprite)

	return nil
}

func (s *SpriteManager) Cleanup() {
	for i := range s.sprites {
		s.sprites[i].Destroy()
	}
}
