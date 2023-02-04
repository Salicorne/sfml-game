package graphics

import (
	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

type Sprite struct {
	sfml_sprite *sfml.Struct_SS_sfSprite
	name        string
	abspos      Vec2
	textureRect Rect
}

func (s *Sprite) GetSfSprite() sfml.Struct_SS_sfSprite {
	return *s.sfml_sprite
}

func (s *Sprite) GetAbsPos() Vec2 {
	return s.abspos
}

func (s *Sprite) SetAbsPos(v Vec2) {
	s.abspos = v
}
