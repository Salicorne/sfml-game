package game

import (
	. "sfml-test/common"
	"sfml-test/graphics"
)

type GameManager struct {
	actors []*Actor
}

func NewGameManager() GameManager {
	return GameManager{actors: []*Actor{}}
}

func (g *GameManager) AddActor(initialPos Vec2, animable graphics.Animable, drawable graphics.Drawable) *Actor {
	a := &Actor{
		abspos:   initialPos,
		animable: animable,
		drawable: drawable,
	}

	g.actors = append(g.actors, a)

	return a
}
