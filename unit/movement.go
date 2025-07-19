package unit

import (
	"image"
	"math"
	"travel-the-world/common"
	"travel-the-world/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

func (u *Unit) Move(objects []*tiles.ObjectTile, units []*Unit, levelDimentions image.Point, cameraPos image.Point) {

	u.updateAngle(levelDimentions, cameraPos)
	u.updateDirection()
	iset := &common.InteractableList{}
	for _, npc := range units {
		iset.Add(npc)
	}
	u.tryMove(iset)

	u.updateAnamationType()

	u.holdUnitInBorderMap(levelDimentions)
}

func (u *Unit) updateAngle(levelDimentions image.Point, cameraPos image.Point) {
	u.vx = 0
	u.vy = 0

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		GoalX, GoalY := ebiten.CursorPosition()
		u.GoalX = float64(GoalX + cameraPos.X)
		u.GoalY = float64(GoalY + cameraPos.Y)
	}

	if u.Goal().In(u.Rect()) {
		u.stopUnit()
	} else {
		goalVector := image.Rectangle{u.Centr(), u.Goal()}
		u.Angle = math.Atan2(float64(goalVector.Dx()), float64(goalVector.Dy()))
		u.countSpeed()
	}
}

func (u *Unit) countSpeed() {
	u.vx = speed * math.Sin(u.Angle)
	u.vy = speed * math.Cos(u.Angle)
}

func (u *Unit) updateDirection() {
	switch {
	case 2.8 <= u.Angle || u.Angle <= -2.8:
		u.animation.direction = DirUp
	case 1.7 <= u.Angle && u.Angle < 2.8:
		u.animation.direction = DirRightUp
	case 1.3 <= u.Angle && u.Angle < 1.7:
		u.animation.direction = DirRight
	case 0.3 <= u.Angle && u.Angle < 1.3:
		u.animation.direction = DirRightDown
	case 0.3 > u.Angle && u.Angle > -0.3:
		u.animation.direction = DirDown
	case -0.3 >= u.Angle && u.Angle > -1.3:
		u.animation.direction = DirLeftDown
	case -1.3 >= u.Angle && u.Angle > -1.7:
		u.animation.direction = DirLeft
	case -1.7 >= u.Angle && u.Angle > -2.8:
		u.animation.direction = DirLeftUp
	default:
		u.animation.direction = DirLeft
	}
}

func (u *Unit) tryMove(objects *common.InteractableList) {
	u.X += u.vx
	u.Y += u.vy
	if objects != nil && u.faceWithObjects(objects) {
		u.X -= u.vx
	}
	if objects != nil && u.faceWithObjects(objects) {
		u.Y -= u.vy
	}
}

func (u *Unit) stopUnit() {
	u.GoalX = float64(u.X)
	u.GoalY = float64(u.Y)
	u.vx = 0
	u.vy = 0
}

func (u *Unit) holdUnitInBorderMap(levelDimentions image.Point) {
	// Не выходим за границы карты

	top := image.Point{1600, 0}
	bottom := image.Point{1600, 1600}
	topToUnitVector := image.Rectangle{top, u.Centr()}
	bottomToUnitVector := image.Rectangle{bottom, u.Centr()}
	u.TopAngle = math.Atan2(float64(topToUnitVector.Dx()), float64(topToUnitVector.Dy()))
	u.BottomAngle = math.Atan2(float64(bottomToUnitVector.Dx()), float64(bottomToUnitVector.Dy()))

	if 0 > u.TopAngle && u.TopAngle < -1.11 || 0 < u.TopAngle && u.TopAngle > 1.11 {
		u.X = 1600
		u.Y = 800
	}

	if 0 > u.BottomAngle && u.BottomAngle > -2.03 || 0 < u.BottomAngle && u.BottomAngle < 2.03 {
		u.X = 1600
		u.Y = 800
	}
}
