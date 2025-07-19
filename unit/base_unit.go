package unit

import (
	"image"
	"travel-the-world/assets"
)

type Unit struct {
	X, Y         float64 // позиция юнита
	vx, vy       float64 // скорость юнита
	GoalX, GoalY float64 // желаемая позиция юнита
	Angle        float64 // угол траектории движения юнита относительно оси y в радианах
	sx, sy       float64 // масштабирование по x и y
	animation    *animation
	TopAngle     float64
	BottomAngle  float64
	Health       int
}

func NewUnit(x, y float64, am *assets.Manager) *Unit {
	a := NewAnimation(x, y, am)

	u := &Unit{
		X:         x,
		Y:         y,
		vx:        0,
		vy:        0,
		Angle:     0,
		sx:        1,
		sy:        1,
		animation: a,
		Health:    100,
	}
	u.GoalX = float64(u.X)
	u.GoalY = float64(u.Y)
	return u
}

func IsoToWorld(isoX, isoY float64, levelDimentions image.Point) (x, y float64) {
	x = (isoX + isoY) - float64(levelDimentions.X/2)
	y = (isoX - isoY)
	return
}

func WorldToIso(x, y float64, levelDimentions image.Point) (isoX, isoY float64) {
	isoX = (x - y) + float64(levelDimentions.X/2)
	isoY = (x + y)
	return
}
