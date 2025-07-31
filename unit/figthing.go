package unit

import (
	"math/rand/v2"
	"travel-the-world/common"
)

func (u *Unit) TakeDamage(dmg int) bool {
	if u.IsAlive() {
		u.playSwordAttack()
		u.Stats.health -= dmg
		u.Stats.stamina -= dmg / 2
		if u.IsDead() {
			playSound(u.sound.diePool, 1)
			u.Command(Die, nil)
		}
		return true
	} else {
		return false
	}
}

func (u *Unit) GetDamage() int {
	kStamina := u.Stats.stamina / u.Stats.maxStamina
	crit := rand.Float64() + float64(kStamina)
	if crit < 1 {
		crit = 1
	}
	return u.Stats.damage * kStamina * int(crit)
}

func (u *Unit) Damage(target common.Damagable) {
	if target != nil {
		if attackedUnit, ok := target.(common.Fightable); ok && target.TakeDamage(u.Stats.damage) {
			if target.IsAlive() {
				u.Stats.stamina--
				attackedUnit.Attack(u)
			} else {
				u.AddExp(10)
				u.Command(Stop, nil)
			}
		}
	}
}

func (u *Unit) Attack(target common.Damagable) {
	if (u.action.command != Attack || u.closerThenCurrentTarget(target)) && target.IsAlive() {
		u.Command(Attack, target)
	}
}

func (u *Unit) IsDead() bool {
	return u.Stats.health <= 0
}

func (u *Unit) IsAlive() bool {
	return u.Stats.health > 0
}

func (u *Unit) isInAttackDistance(obj common.Target) bool {
	distance := u.DistanceTo(obj.Point())
	return distance <= u.Stats.AttackDistance
}
