package unit

import (
	"travel-the-world/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

func (u *Unit) updateAction() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		u.stopUnit()
		u.SetAction(ActionAttack)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		u.stopUnit()
		u.SetAction(ActionShoot)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		u.stopUnit()
		u.SetAction(ActionDie)
	} else if u.vx == 0 || u.vy == 0 {
		u.SetAction(ActionIdle)
	} else {
		u.SetAction(ActionRun)
	}
}

func (u *Unit) SetAction(newAction assets.Action) {
	if u.Action != newAction {
		u.current = 0
		u.Action = newAction
	}
}
