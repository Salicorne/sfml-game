package graphics

import (
	"time"

	. "github.com/Salicorne/sfml-game/common"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

type Drawable interface {
	Draw(sfml.Struct_SS_sfRenderWindow, Vec2)
	SetAbsPos(Vec2)
	GetAbsPos() Vec2
	Destroy()
}

type Animable interface {
	Animate(elapsed time.Duration)
	PlayAnimation(animation AnimationType)
}
