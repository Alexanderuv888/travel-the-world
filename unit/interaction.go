package unit

import (
	"image"
	"math"
	"travel-the-world/common"
)

func (u *Unit) faceWithObjects(objects *common.InteractableList) bool {
	for _, o := range objects.Items {
		if u != o && u.Rect().Overlaps(o.Rect()) {
			if target, ok := o.(common.Damagable); ok {
				if target.IsDead() {
					return false
				}
			}
			overlapSize := u.Point().Sub(o.Point())
			u.X = float64(u.Point().X + overlapSize.X/4)
			u.Y = float64(u.Point().Y + overlapSize.Y/4)
			return true
		}
	}
	return false
}

func (u *Unit) ObjType() common.ObjType {
	return common.Unit
}

func (u *Unit) Interact(obj common.Interactable) {
	if u.Rect().Overlaps(obj.Rect()) {
		obj.Rect().Sub(u.Point())
	}
}

func (u *Unit) InteractWithAll(objects *common.InteractableList) {
	if u.IsDead() {
		return
	}
	for _, obj := range objects.Items {
		u.interactWith(obj)
	}

	u.Walk()
}

func (u *Unit) closerThenCurrentTarget(target common.Target) bool {
	if u.target == nil {
		return true
	}
	newTargetDistance := u.DistanceTo(target.Point())
	currentTargetDistance := u.DistanceTo(u.target.Point())

	return newTargetDistance < currentTargetDistance
}

func (u *Unit) IsInVision(obj common.Interactable) bool {
	return u.DistanceTo(obj.Point()) <= u.Stats.Vision
}

func (u *Unit) DistanceTo(point image.Point) float64 {
	dx := float64(point.X - u.Point().X)
	dy := float64(point.Y - u.Point().Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func (u Unit) IsEnemy(unit common.Damagable) bool {
	return true
}

func (u *Unit) catchTarget(obj common.Interactable) {
	if target, ok := obj.(common.Damagable); ok {
		if target.IsAlive() {
			u.Command(MoveTo, target)
		}
	}
}

func (u *Unit) Rect() image.Rectangle {
	return image.Rect(int(u.X-(frameWidth/8)), int(u.Y-(frameHeight/8)), int(u.X+(frameWidth/8)), int(u.Y+(frameHeight/8)))
}

func centrOfRect(r image.Rectangle) image.Point {
	return image.Point{(r.Max.X + r.Min.X) / 2, (r.Max.Y + r.Min.Y) / 2}
}

func (u *Unit) Point() image.Point {
	return image.Point{int(u.X), int(u.Y)}
}
