package unit

import (
	"image"
	"math"
	"math/rand"
	"time"
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
		if u == obj {
			continue
		}
		if u.isInVision(obj) {
			if target, ok := obj.(common.Damagable); ok {
				if target.IsAlive() && u.closerThenCurrentTarget(target) {
					u.Command(Attack, target)
				}
			}
		}
	}

	if u.target == nil {
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		x := r.Intn(3200)
		y := r.Intn(1600)
		u.MoveToPoint(image.Point{x, y})
	}
}

func (u *Unit) closerThenCurrentTarget(target common.Target) bool {
	if u.target == nil {
		return true
	}
	dx := float64(target.Point().X - u.Point().X)
	dy := float64(target.Point().Y - u.Point().Y)
	newTargetDistance := math.Sqrt(dx*dx + dy*dy)

	cdx := float64(u.target.Point().X - u.Point().X)
	cdy := float64(u.target.Point().Y - u.Point().Y)
	currentTargetDistance := math.Sqrt(cdx*cdx + cdy*cdy)

	return newTargetDistance <= currentTargetDistance
}

func (u *Unit) isInVision(obj common.Interactable) bool {
	dx := float64(obj.Point().X - u.Point().X)
	dy := float64(obj.Point().Y - u.Point().Y)
	distance := math.Sqrt(dx*dx + dy*dy)

	return distance <= u.stats.vision
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
