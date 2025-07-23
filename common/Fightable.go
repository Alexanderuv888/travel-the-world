package common

type Fightable interface {
	Attack(d *Damagable)
}
