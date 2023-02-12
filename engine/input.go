package engine

import (
	"math"
	. "sfml-test/common"

	sfmlwindow "github.com/telroshan/go-sfml/v2/window"
)

func GetInput() Vec2 {
	m := Vec2{}
	inputIsNull := true
	if sfmlwindow.SfKeyboard_isKeyPressed(sfmlwindow.SfKeyCode(sfmlwindow.SfKeyZ)) > 0 {
		m.Y--
		inputIsNull = false
	}
	if sfmlwindow.SfKeyboard_isKeyPressed(sfmlwindow.SfKeyCode(sfmlwindow.SfKeyQ)) > 0 {
		m.X--
		inputIsNull = false
	}
	if sfmlwindow.SfKeyboard_isKeyPressed(sfmlwindow.SfKeyCode(sfmlwindow.SfKeyS)) > 0 {
		m.Y++
		inputIsNull = false
	}
	if sfmlwindow.SfKeyboard_isKeyPressed(sfmlwindow.SfKeyCode(sfmlwindow.SfKeyD)) > 0 {
		m.X++
		inputIsNull = false
	}
	if inputIsNull || (m.X == 0 && m.Y == 0) {
		return Vec2{X: 0, Y: 0}
	}
	l := float32(math.Sqrt(float64(m.X*m.X + m.Y*m.Y)))

	return Vec2{X: m.X / l, Y: m.Y / l}
}
