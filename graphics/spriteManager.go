package graphics

import (
	"fmt"
	"sort"
	"time"

	. "sfml-test/common"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

type SpriteManager struct {
	txtrmgr   TextureManager
	sprites   []Drawable
	animables []Animable
}

func NewSpriteManager(txtrmgr TextureManager) SpriteManager {
	return SpriteManager{
		txtrmgr: txtrmgr,
		sprites: []Drawable{},
	}
}

func (s *SpriteManager) Draw(w sfml.Struct_SS_sfRenderWindow, winpos Vec2) {
	sort.SliceStable(s.sprites, func(i, j int) bool { return s.sprites[i].GetAbsPos().Y < s.sprites[j].GetAbsPos().Y })
	for i := range s.sprites {
		s.sprites[i].Draw(w, winpos)
	}
}

func (s *SpriteManager) Animate(elapsed time.Duration) {
	for i := range s.animables {
		s.animables[i].Animate(elapsed)
	}
}

func (s *SpriteManager) LoadBasicSprite(spriteName, textureName string, textureRect Rect, feetPos Vec2, initialPos Vec2) (*BasicSprite, error) {
	txtr, err := s.txtrmgr.GetSfTexture(textureName)
	if err != nil {
		return nil, err
	}
	spr := sfml.SfSprite_create()
	sfml.SfSprite_setTexture(spr, *txtr, 1)

	sfml.SfSprite_setTextureRect(spr, textureRect.ToSFMLIntRect())

	rect := sfml.SfRectangleShape_create()
	sfml.SfRectangleShape_setSize(rect, Vec2{X: 6, Y: 6}.ToSFMLVector2f())
	sfml.SfRectangleShape_setFillColor(rect, sfml.GetSfBlue())

	sprite := &BasicSprite{sfml_sprite: &spr, name: spriteName, feetPos: feetPos, nextabspos: initialPos, abspos: initialPos, sfml_rect: &rect}
	s.sprites = append(s.sprites, sprite)

	return sprite, nil
}

func (s *SpriteManager) LoadAnimatedSprite(spriteName, textureName string, playbackMode PlaybackMode, animations map[AnimationType]Animation, initialAnimation AnimationType, feetPos Vec2, initialPos Vec2) (*AnimatedSprite, error) {
	if len(animations) == 0 {
		return nil, fmt.Errorf("Animated sprites must have at least one animation")
	}
	if _, ok := animations[initialAnimation]; !ok {
		return nil, fmt.Errorf("Actor loaded without its initial animations")
	}

	txtr, err := s.txtrmgr.GetSfTexture(textureName)
	if err != nil {
		return nil, err
	}
	spr := sfml.SfSprite_create()
	sfml.SfSprite_setTexture(spr, *txtr, 1)

	sfml.SfSprite_setTextureRect(spr, animations[initialAnimation][0].Rect.ToSFMLIntRect())

	rect := sfml.SfRectangleShape_create()
	sfml.SfRectangleShape_setSize(rect, Vec2{X: 6, Y: 6}.ToSFMLVector2f())
	sfml.SfRectangleShape_setFillColor(rect, sfml.GetSfGreen())

	sprite := AnimatedSprite{BasicSprite: BasicSprite{sfml_sprite: &spr, sfml_rect: &rect, name: spriteName, feetPos: feetPos, abspos: initialPos, nextabspos: initialPos}, animations: animations, playbackMode: playbackMode, currentAnimation: initialAnimation, nextAnimation: initialAnimation, currentFrame: 0}
	s.sprites = append(s.sprites, &sprite)
	s.animables = append(s.animables, &sprite)

	return &sprite, nil
}

func (s *SpriteManager) Cleanup() {
	for i := range s.sprites {
		s.sprites[i].Destroy()
	}
}
