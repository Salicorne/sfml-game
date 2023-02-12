package game

import (
	. "github.com/Salicorne/sfml-game/common"
	"github.com/Salicorne/sfml-game/graphics"
)

type Actor struct {
	abspos    Vec2
	animable  graphics.Animable
	drawable  graphics.Drawable
	direction Direction
	isMoving  bool
}

func (a *Actor) Move(delta Vec2) {
	a.abspos.X += delta.X
	a.abspos.Y += delta.Y
	if delta.Len() > 0.1 {
		a.direction = Vec2ToDirection(delta)
		a.isMoving = true
	} else {
		a.isMoving = false
	}
	a.animable.PlayAnimation(DirectionToAnimationType(a.direction, a.isMoving))

	a.drawable.SetAbsPos(a.abspos)
}
