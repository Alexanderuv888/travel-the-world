package unit

import (
	"image"
	"travel-the-world/assets"
	"travel-the-world/common"
	"travel-the-world/world"
)

type Unit struct {
	am          *assets.Manager
	sound       *sound
	Stats       *Stats
	X, Y        float64       // позиция юнита
	vx, vy      float64       // скорость юнита
	target      common.Target // желаемая позиция юнита
	Angle       float64       // угол траектории движения юнита относительно оси y в радианах
	sx, sy      float64       // масштабирование по x и y
	action      action
	TopAngle    float64
	BottomAngle float64
	Behavior    Behavior
	highlight   int
}

func NewUnit(x, y float64, am *assets.Manager, behavior Behavior) *Unit {
	a := NewAnimation(x, y, am)
	action := action{Stop, nil, a, false}
	stats := NewStats(12, 12, 2, 50, 350, speed)

	u := &Unit{
		am:       am,
		sound:    NewSound(am),
		Stats:    stats,
		X:        x,
		Y:        y,
		vx:       0,
		vy:       0,
		Angle:    0,
		sx:       1,
		sy:       1,
		action:   action,
		Behavior: behavior,
	}
	return u
}

func (u *Unit) Update(worldCtx *world.Context, levelDimentions image.Point) {

	u.updateAction(&worldCtx.InteractableList, levelDimentions)
	if u.IsAlive() {
		u.Move(worldCtx)
		u.Behavior.Update(u, worldCtx)
	}
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

func (u *Unit) GoalX() float64 {
	if u.target != nil {
		return float64(u.target.Point().X)
	}
	return float64(u.Point().X)
}

func (u *Unit) GoalY() float64 {
	if u.target != nil {
		return float64(u.target.Point().Y)
	}
	return float64(u.Point().Y)
}
