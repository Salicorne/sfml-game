package graphics

import (
	"fmt"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

type SpriteManager struct {
	txtrmgr TextureManager
	sprites []Sprite
}

func NexSpriteManager(txtrmgr TextureManager) SpriteManager {
	return SpriteManager{
		txtrmgr: txtrmgr,
		sprites: []Sprite{},
	}
}

func (s *SpriteManager) LoadSprite(spriteName, textureName string, textureRect Rect) error {
	txtr, err := s.txtrmgr.GetSfTexture(textureName)
	if err != nil {
		return err
	}
	spr := sfml.SfSprite_create()
	sfml.SfSprite_setTexture(spr, *txtr, 1)

	sfml.SfSprite_setTextureRect(spr, textureRect.ToSFMLIntRect())

	s.sprites = append(s.sprites, Sprite{sfml_sprite: &spr, name: spriteName, abspos: Vec2{0, 0}})

	return nil
}

func (s *SpriteManager) GetSprite(spriteName string) (*Sprite, error) {
	for i := range s.sprites {
		if s.sprites[i].name == spriteName {
			return &s.sprites[i], nil
		}
	}
	return nil, fmt.Errorf("No sprite is referenced as %s", spriteName)
}

func (s *SpriteManager) Cleanup() {
	for i := range s.sprites {
		sfml.SfSprite_destroy(*s.sprites[i].sfml_sprite)
	}
}
