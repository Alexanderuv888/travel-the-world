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
		u.setAnimation(ActionAttack, true)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
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
		u.moveTo(iset, levelDimentions)
	case Attack:
		u.Attack(iset, levelDimentions)
	case Die:
		u.Die()
	default:
		u.setAnimation(ActionIdle, true)
		u.stopUnit()
	}

}

func (u *Unit) moveTo(iset *common.InteractableList, levelDimentions image.Point) {
	u.setAnimation(ActionRun, false)
	if u.target == nil || u.target.Point().In(u.Rect()) {
		u.Command(Stop, nil)
	} else {
		u.Move(iset, levelDimentions)
	}
}

func (u *Unit) Attack(iset *common.InteractableList, levelDimentions image.Point) {
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
			u.moveTo(iset, levelDimentions)
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
	u.stats.status = Dead
	if u.action.isFinished() {
		u.action.animation.freez = true
	}
}
