package unit

import (
	"travel-the-world/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

func (u *Unit) updateAnamation() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		u.stopUnit()
		u.setAnimation(ActionAttack)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		u.stopUnit()
		u.setAnimation(ActionShoot)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		u.stopUnit()
		u.setAnimation(ActionDie)
	} else if u.vx == 0 || u.vy == 0 {
		u.setAnimation(ActionIdle)
	} else {
		u.setAnimation(ActionRun)
	}
}

func (u *Unit) setAnimation(newAction assets.Action) {
	if u.Action != newAction {
		u.current = 0
		u.Action = newAction
	}
}
