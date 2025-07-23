package unit

import (
	"image"
	"travel-the-world/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Command string

const (
	MoveTo Command = "moveTo"
	Attack Command = "attack"
	Stop   Command = "stop"
	Follow Command = "follow"
	Die    Command = "die"
)

func (u *Unit) Command(command Command, target common.Target) {
	u.action.command = command
	u.target = target
}

func (u *Unit) ListenKeyBoard(cameraPos image.Point) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		GoalX, GoalY := ebiten.CursorPosition()
		u.MoveToPoint(image.Point{GoalX + cameraPos.X, GoalY + cameraPos.Y})
	}
}

func (u *Unit) MoveToPoint(p image.Point) {
	target := common.TargetFunc(func() image.Point {
		return p
	})
	u.Command(MoveTo, target)
}
