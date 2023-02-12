package graphics

import (
	. "sfml-test/common"
	"time"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

/**************************
*       BasicSprite       *
**************************/

type BasicSprite struct {
	sfml_sprite *sfml.Struct_SS_sfSprite
	sfml_rect   *sfml.Struct_SS_sfRectangleShape
	name        string
	abspos      Vec2
	nextabspos  Vec2
	textureRect Rect
	nextRect    Rect
	currectRect Rect
}

func (s *BasicSprite) Draw(w sfml.Struct_SS_sfRenderWindow, winpos Vec2) {
	s.abspos = s.nextabspos

	if s.currectRect != s.nextRect {
		s.currectRect = s.nextRect
		sfml.SfSprite_setTextureRect(s.GetSfSprite(), s.currectRect.ToSFMLIntRect())
	}
	sfml.SfSprite_setPosition(*s.sfml_sprite, Vec2{X: s.abspos.X - winpos.X, Y: s.abspos.Y - winpos.Y}.ToSFMLVector2f())
	sfml.SfRectangleShape_setPosition(*s.sfml_rect, Vec2{X: s.abspos.X - winpos.X, Y: s.abspos.Y - winpos.Y}.ToSFMLVector2f())

	sfml.SfRenderWindow_drawSprite(w, *s.sfml_sprite, (sfml.SfRenderStates)(sfml.SwigcptrSfRenderStates(0)))
	sfml.SfRenderWindow_drawRectangleShape(w, *s.sfml_rect, (sfml.SfRenderStates)(sfml.SwigcptrSfRenderStates(0)))
	//fmt.Printf("%d %d %d %d\n", s.currectRect.X, s.currectRect.Y, s.currectRect.W, s.currectRect.H)
	//fmt.Printf("%f %f\n", s.abspos.X, s.abspos.Y)
}

func (s *BasicSprite) Destroy() {
	sfml.SfSprite_destroy(*s.sfml_sprite)
	sfml.SfRectangleShape_destroy(*s.sfml_rect)
}

func (s *BasicSprite) GetSfSprite() sfml.Struct_SS_sfSprite {
	return *s.sfml_sprite
}

func (s *BasicSprite) GetAbsPos() Vec2 {
	return s.nextabspos
}

func (s *BasicSprite) SetAbsPos(v Vec2) {
	s.nextabspos = v
}

/***************************
*      AnimatedSprite      *
***************************/

type AnimatedSpriteFrame struct {
	Rect     Rect
	Duration time.Duration
}

type Animation []AnimatedSpriteFrame

type AnimatedSprite struct {
	BasicSprite
	animations       map[AnimationType]Animation
	currentAnimation AnimationType
	nextAnimation    AnimationType
	counter          time.Duration
	currentFrame     int
	playbackMode     PlaybackMode
}

func (s *AnimatedSprite) Animate(elapsed time.Duration) {
	s.counter += elapsed

	if s.currentAnimation != s.nextAnimation {
		s.currentFrame = min(s.currentFrame, len(s.animations[s.nextAnimation])-1)
		s.nextRect = s.animations[s.nextAnimation][s.currentFrame].Rect
		s.currentAnimation = s.nextAnimation
	}

	for s.counter >= s.animations[s.currentAnimation][s.currentFrame].Duration {
		s.counter -= s.animations[s.currentAnimation][s.currentFrame].Duration

		if s.currentFrame+1 < len(s.animations[s.currentAnimation]) {
			s.currentFrame += 1
		} else if s.playbackMode == PlaybackMode_LOOP {
			s.currentFrame = 0
		}

	}
	s.nextRect = s.animations[s.nextAnimation][s.currentFrame].Rect
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (s *AnimatedSprite) PlayAnimation(animation AnimationType) {
	if _, ok := s.animations[animation]; !ok {
		Error("Animated sprite %s has no animation %s", s.name, animation)
	} else {
		s.nextAnimation = animation
	}
}
