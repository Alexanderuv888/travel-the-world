package unit

import (
	"image"
	"travel-the-world/common"
)

type action struct {
	command   Command
	target    *common.Interactable
	animation *animation
	finished  bool
}

func (u *Unit) updateAction(iset *common.InteractableList, levelDimentions image.Point) {
	/*if ebiten.IsKeyPressed(ebiten.KeyA) {
		u.target = u
		u.Attack(iset, levelDimentions)
	}
	else if ebiten.IsKeyPressed(ebiten.KeyS) {
		u.setAnimation(ActionShoot, true)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		u.setAnimation(ActionDie, true)
	} else if u.vx == 0 || u.vy == 0 {
		u.setAnimation(ActionIdle, true)
	} else {
		u.setAnimation(ActionRun, false)
	}*/

	switch u.action.command {
	case MoveTo:
		u.moveToTarget()
	case Attack:
		u.attack()
	case Die:
		u.Die()
	default:
		u.setAnimation(ActionIdle, true)
		u.stopUnit()
	}

}

func (u *Unit) moveToTarget() {
	u.setAnimation(ActionRun, false)
	if u.target == nil || u.isInAttackDistance(u.target) {
		u.Command(Stop, nil)
	}
}

func (u *Unit) attack() {
	if target, ok := u.target.(common.Damagable); ok {
		if target.IsDead() {
			u.Command(Stop, nil)
			return
		}
		if u.isInAttackDistance(target) {
			u.updateAngle()
			u.setAnimation(ActionAttack, true)
			if u.action.isFinished() {
				u.Damage(target)
				u.action.finished = true
			}
		} else {
			u.moveToTarget()
		}
	} else {
		u.Command(Stop, nil)
	}
}

func (a *action) isFinished() bool {
	if (a.animation.current >= a.animation.len-1) && !a.finished {
		a.finished = true
		return true
	}
	if !(a.animation.current >= a.animation.len-1) && a.finished {
		a.finished = false
	}
	return false
}

func (u *Unit) Die() {
	u.setAnimation(ActionDie, true)
	if u.action.isFinished() {
		u.action.animation.freez = true
	}
}
