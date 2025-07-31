package unit

import (
	"image"
	"math"
	"travel-the-world/common"
	"travel-the-world/world"
)

func (u *Unit) Move(worldCtx *world.Context) {
	if u.target == nil {
		return
	}
	u.updateAngle()
	u.updateDirection()
	u.tryMove(&worldCtx.InteractableList)

	if reachBorderMap(u.Point()) {
		u.Command(Stop, nil)
	}
}

func (u *Unit) updateAngle() {
	u.vx = 0
	u.vy = 0

	if u.target != nil {
		u.Angle = countAngle(u.Point(), u.target.Point())
		u.countSpeed()
	}
}

func countAngle(p1 image.Point, p2 image.Point) float64 {
	goalVector := image.Rectangle{p1, p2}
	return math.Atan2(float64(goalVector.Dx()), float64(goalVector.Dy()))
}

func (u *Unit) countSpeed() {
	u.vx = speed * math.Sin(u.Angle)
	u.vy = speed * math.Cos(u.Angle)
}

func (u *Unit) updateDirection() {
	switch {
	case 2.8 <= u.Angle || u.Angle <= -2.8:
		u.action.animation.direction = DirUp
	case 1.7 <= u.Angle && u.Angle < 2.8:
		u.action.animation.direction = DirRightUp
	case 1.3 <= u.Angle && u.Angle < 1.7:
		u.action.animation.direction = DirRight
	case 0.3 <= u.Angle && u.Angle < 1.3:
		u.action.animation.direction = DirRightDown
	case 0.3 > u.Angle && u.Angle > -0.3:
		u.action.animation.direction = DirDown
	case -0.3 >= u.Angle && u.Angle > -1.3:
		u.action.animation.direction = DirLeftDown
	case -1.3 >= u.Angle && u.Angle > -1.7:
		u.action.animation.direction = DirLeft
	case -1.7 >= u.Angle && u.Angle > -2.8:
		u.action.animation.direction = DirLeftUp
	default:
		u.action.animation.direction = DirLeft
	}
}

func (u *Unit) tryMove(objects *common.InteractableList) {
	if u.isInAttackDistance(u.target) {
		return
	}
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
	u.vx = 0
	u.vy = 0
}

func reachBorderMap(p image.Point) bool {
	// Не выходим за границы карты
	top := image.Point{1600, 0}
	left := image.Point{0, 800}
	bottom := image.Point{1600, 1600}
	right := image.Point{3200, 800}
	topToLeftVector := image.Rectangle{top, left}
	topToRightVector := image.Rectangle{top, right}
	bottomToLeftVector := image.Rectangle{bottom, left}
	bottomToRightVector := image.Rectangle{bottom, right}
	topToUnitVector := image.Rectangle{top, p}
	bottomToUnitVector := image.Rectangle{bottom, p}

	topLeftAngle := math.Atan2(float64(topToLeftVector.Dx()), float64(topToLeftVector.Dy()))
	topRightAngle := math.Atan2(float64(topToRightVector.Dx()), float64(topToRightVector.Dy()))
	bottomLeftAngle := math.Atan2(float64(bottomToLeftVector.Dx()), float64(bottomToLeftVector.Dy()))
	bottomRightAngle := math.Atan2(float64(bottomToRightVector.Dx()), float64(bottomToRightVector.Dy()))
	/*topLeftAngle = -1.11
	topRightAngle = 1.11
	bottomLeftAngle = -2.03
	bottomRightAngle = 2.03*/
	topAngle := math.Atan2(float64(topToUnitVector.Dx()), float64(topToUnitVector.Dy()))
	bottomAngle := math.Atan2(float64(bottomToUnitVector.Dx()), float64(bottomToUnitVector.Dy()))

	if 0 > topAngle && topAngle < topLeftAngle || 0 < topAngle && topAngle > topRightAngle {
		return true
	}

	if 0 > bottomAngle && bottomAngle > bottomLeftAngle || 0 < bottomAngle && bottomAngle < bottomRightAngle {
		return true
	}
	return false
}
