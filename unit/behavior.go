package unit

import (
	"image"
	"math/rand"
	"time"
	"travel-the-world/common"
)

func (u *Unit) interactWith(obj common.Interactable) {
	if u == obj {
		return
	}
	if u.isInVision(obj) {
		if target, ok := obj.(common.Damagable); ok {
			if target.IsAlive() && u.closerThenCurrentTarget(target) {
				u.Command(Attack, target)
			}
		}
	}
}

func (u *Unit) walk() {
	if u.target == nil {
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		x := r.Intn(3200)
		y := r.Intn(1600)
		u.MoveToPoint(image.Point{x, y})
	}
}
