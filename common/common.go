package common

import (
	"fmt"
	"math"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

type PlaybackMode int

const (
	PlaybackMode_LOOP   PlaybackMode = iota
	PlaybackMode_SINGLE PlaybackMode = iota
)

type AnimationType int

const (
	AnimationType_IDLEN AnimationType = iota
	AnimationType_IDLES AnimationType = iota
	AnimationType_IDLEE AnimationType = iota
	AnimationType_IDLEW AnimationType = iota
	AnimationType_WALKN AnimationType = iota
	AnimationType_WALKS AnimationType = iota
	AnimationType_WALKE AnimationType = iota
	AnimationType_WALKW AnimationType = iota
)

type Direction int

const (
	Direction_N Direction = iota
	Direction_W Direction = iota
	Direction_S Direction = iota
	Direction_E Direction = iota
)

func Vec2ToDirection(v Vec2) Direction {
	h := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))

	if 2*v.X >= math.Sqrt2*h {
		return Direction_E
	} else if -2*v.X >= math.Sqrt2*h {
		return Direction_W
	} else if 2*v.Y <= math.Sqrt2*h {
		return Direction_N
	}
	return Direction_S
}

var dirToAnimWalking = map[Direction]AnimationType{
	Direction_N: AnimationType_WALKN,
	Direction_E: AnimationType_WALKE,
	Direction_S: AnimationType_WALKS,
	Direction_W: AnimationType_WALKW,
}
var dirToAnimIdle = map[Direction]AnimationType{
	Direction_N: AnimationType_IDLEN,
	Direction_E: AnimationType_IDLEE,
	Direction_S: AnimationType_IDLES,
	Direction_W: AnimationType_IDLEW,
}

func DirectionToAnimationType(d Direction, moving bool) AnimationType {
	if moving {
		return dirToAnimWalking[d]
	}
	return dirToAnimIdle[d]
}

type Vec2 struct {
	X float32
	Y float32
}

func (v Vec2) ToSFMLVector2f() sfml.SfVector2f {
	vec := sfml.NewSfVector2f()
	vec.SetX(v.X)
	vec.SetY(v.Y)
	return vec
}

func (v Vec2) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
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

func Error(format string, args ...interface{}) {
	fmt.Printf("[ERR] "+format+"\n", args...)
}
