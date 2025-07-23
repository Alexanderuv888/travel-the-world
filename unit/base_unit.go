package unit

import (
	"image"
	"travel-the-world/assets"
	"travel-the-world/common"
	"travel-the-world/tiles"
)

type Unit struct {
	am          *assets.Manager
	stats       *stats
	X, Y        float64       // позиция юнита
	vx, vy      float64       // скорость юнита
	target      common.Target // желаемая позиция юнита
	Angle       float64       // угол траектории движения юнита относительно оси y в радианах
	sx, sy      float64       // масштабирование по x и y
	action      action
	TopAngle    float64
	BottomAngle float64
}

func NewUnit(x, y float64, am *assets.Manager) *Unit {
	a := NewAnimation(x, y, am)
	action := action{Stop, nil, a, false}
	stats := stats{Alive, 12, 12, 2, 50, 350, speed}

	u := &Unit{
		am:     am,
		stats:  &stats,
		X:      x,
		Y:      y,
		vx:     0,
		vy:     0,
		Angle:  0,
		sx:     1,
		sy:     1,
		action: action,
	}
	return u
}

func (u *Unit) Update(objects []*tiles.ObjectTile, units []*Unit, levelDimentions image.Point) {
	iset := &common.InteractableList{}
	for _, npc := range units {
		iset.Add(npc)
	}
	u.updateAction(iset, levelDimentions)

	u.InteractWithAll(iset)
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
