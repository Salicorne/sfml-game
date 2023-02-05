package graphics

import (
	"time"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

type Drawable interface {
	Draw(sfml.Struct_SS_sfRenderWindow)
	Destroy()
}

type Animable interface {
	Animate(elapsed time.Duration)
}

type Vec2 struct {
	X float32
	Y float32
}

type Rect struct {
	X int
	Y int
	W int
	H int
}

func (r Rect) ToSFMLIntRect() sfml.SfIntRect {
	rect := sfml.NewSfIntRect()
	rect.SetLeft(r.X)
	rect.SetTop(r.Y)
	rect.SetWidth(r.W)
	rect.SetHeight(r.H)

	return rect
}
