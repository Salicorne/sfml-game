package graphics

import (
	"time"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

/**************************
*       BasicSprite       *
**************************/

type BasicSprite struct {
	sfml_sprite *sfml.Struct_SS_sfSprite
	name        string
	abspos      Vec2
	textureRect Rect
}

func (s BasicSprite) Draw(w sfml.Struct_SS_sfRenderWindow) {
	sfml.SfRenderWindow_drawSprite(w, *s.sfml_sprite, (sfml.SfRenderStates)(sfml.SwigcptrSfRenderStates(0)))
}

func (s BasicSprite) Destroy() {
	sfml.SfSprite_destroy(*s.sfml_sprite)
}

func (s *BasicSprite) GetSfSprite() sfml.Struct_SS_sfSprite {
	return *s.sfml_sprite
}

func (s *BasicSprite) GetAbsPos() Vec2 {
	return s.abspos
}

func (s *BasicSprite) SetAbsPos(v Vec2) {
	s.abspos = v
}

/***************************
*      AnimatedSprite      *
***************************/

type PlaybackMode int

const (
	PlaybackMode_LOOP   PlaybackMode = iota
	PlaybackMode_SINGLE PlaybackMode = iota
)

type AnimatedSpriteFrame struct {
	Rect     Rect
	Duration time.Duration
}

type AnimatedSprite struct {
	BasicSprite
	frames       []AnimatedSpriteFrame
	counter      time.Duration
	currentFrame int
	playbackMode PlaybackMode
}

func (s *AnimatedSprite) Animate(elapsed time.Duration) {
	s.counter += elapsed
	for s.counter >= s.frames[s.currentFrame].Duration {
		s.counter -= s.frames[s.currentFrame].Duration

		if s.currentFrame+1 < len(s.frames) {
			s.currentFrame += 1
		} else if s.playbackMode == PlaybackMode_LOOP {
			s.currentFrame = 0
		}

		sfml.SfSprite_setTextureRect(s.BasicSprite.GetSfSprite(), s.frames[s.currentFrame].Rect.ToSFMLIntRect())
	}
}
