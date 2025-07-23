package common

type Damagable interface {
	Target
	TakeDamage(dmg int) bool
	IsAlive() bool
	IsDead() bool
}
