package graphics

import (
	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

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
