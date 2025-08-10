package unit

import (
	"image"
	"math/rand"
	"time"
	"travel-the-world/common"
	"travel-the-world/world"
)

type Behavior interface {
	Update(u *Unit, worldCtx *world.Context)
}

type PlayerBehavior struct{}

func (b *PlayerBehavior) Update(u *Unit, worldCtx *world.Context) {
	if u.ListenKeyBoard(worldCtx.Camera) {
		for _, obj := range worldCtx.InteractableList.Items {
			if u.target.Point().In(obj.Rect()) {
				if target, ok := obj.(common.Damagable); ok && target.IsAlive() && u.IsEnemy(target) {
					u.Command(Attack, target)
				}
				u.target = obj
			}
		}
	}
}

type IdleBehavior struct {
	tick int
}

func (b *IdleBehavior) Update(u *Unit, worldCtx *world.Context) {
	b.tick++
	if b.tick >= 300 {
		u.Walk()
		b.tick = 0
	}
}

type AggressiveBehavior struct{}

func (b *AggressiveBehavior) Update(u *Unit, worldCtx *world.Context) {
	target := findClosestEnemy(u, worldCtx.InteractableList)
	if target != nil {
		u.Command(Attack, *target)
	}
}

type PatrolBehavior struct{}

func (b *PatrolBehavior) Update(u *Unit, worldCtx *world.Context) {
	u.Walk()
}

func findClosestEnemy(u *Unit, iList common.InteractableList) *common.Damagable {
	var closest *common.Damagable
	minDist := u.Stats.Vision
	for _, obj := range iList.Items {
		if u == obj {
			continue
		}
		if u.IsInVision(obj) {
			if target, ok := obj.(common.Damagable); ok && target.IsAlive() && u.IsEnemy(target) {
				targetDistance := u.DistanceTo(target.Point())
				if targetDistance < minDist {
					minDist = targetDistance
					closest = &target
				}
			}
		}
	}
	return closest
}

func (u *Unit) interactWith(obj common.Interactable) {
	if u == obj {
		return
	}
	if u.IsInVision(obj) {
		if target, ok := obj.(common.Damagable); ok {
			if target.IsAlive() && u.closerThenCurrentTarget(target) {
				u.Command(Attack, target)
			}
		}
	}
}

func (u *Unit) Walk() {
	if u.target == nil {
		point := u.Point().Add(getRandPointLessThen(int(u.Stats.Vision)))
		/*angle := countAngle(u.Point(), point)
		tx := int(u.Stats.Vision * math.Sin(angle))
		ty := int(u.Stats.Vision * math.Cos(angle))*/
		u.MoveToPoint(point)
	}
}

func getRandPointLessThen(max int) image.Point {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	x := r.Intn(max) - max/2
	y := r.Intn(max) - max/2
	return image.Point{x, y}
}
