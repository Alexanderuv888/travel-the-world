package unit

import (
	"math"
	"travel-the-world/common"
)

func (u *Unit) TakeDamage(dmg int) bool {
	if u.stats.health > 0 {
		u.stats.health -= dmg
		return true
	} else {
		u.Command(Die, nil)
		return false
	}
}

func (u *Unit) Damage(target common.Damagable) {
	target.TakeDamage(u.stats.damage)
}

func (u *Unit) IsDead() bool {
	return u.stats.status == Dead
}

func (u *Unit) IsAlive() bool {
	return u.stats.status != Dead
}

func (u *Unit) isInAttackDistance(obj common.Target) bool {
	dx := float64(obj.Point().X - u.Point().X)
	dy := float64(obj.Point().Y - u.Point().Y)
	distance := math.Sqrt(dx*dx + dy*dy)

	return distance <= u.stats.attackDistance
}
