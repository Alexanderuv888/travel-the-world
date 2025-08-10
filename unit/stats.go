package unit

type Stats struct {
	health              int
	maxHealth           int
	healthSpeedRecover  int
	stamina             int
	maxStamina          int
	staminaSpeedRecover int
	damage              int
	AttackDistance      float64
	Vision              float64
	speed               float64
	level               int
	maxExp              int
	exp                 int
}

func NewStats(health, maxHealth, damage int, attackDisance, vision, speed float64) *Stats {
	return &Stats{
		health:              health,
		maxHealth:           maxHealth,
		healthSpeedRecover:  1000,
		stamina:             health / 2,
		maxStamina:          maxHealth / 2,
		staminaSpeedRecover: 1000,
		damage:              damage,
		AttackDistance:      attackDisance,
		Vision:              vision,
		speed:               speed,
		level:               1,
		maxExp:              10,
		exp:                 0}
}

func (u *Unit) helthLeftInPercent() float64 {
	return float64(u.Stats.health) / float64(u.Stats.maxHealth)
}

func (u *Unit) AddExp(exp int) {
	s := u.Stats
	u.Stats.exp += exp
	if u.Stats.exp >= u.Stats.maxExp {
		expLeft := u.Stats.exp - u.Stats.maxExp
		u.levelUp()
		u.Stats = s
		u.AddExp(expLeft)
	}
}

func (u *Unit) levelUp() {
	u.Stats.level++
	u.Stats.exp = 0
	u.Stats.maxExp = u.Stats.maxExp + factorial(u.Stats.level)
	u.Stats.maxHealth += 2
	u.Stats.health = u.Stats.maxHealth
	u.Stats.maxStamina += 1
	u.Stats.stamina = u.Stats.maxStamina
	u.Stats.damage += 1
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func (u *Unit) Health() int        { return u.Stats.health }
func (u *Unit) MaxHealth() int     { return u.Stats.maxHealth }
func (u *Unit) Stamina() int       { return u.Stats.stamina }
func (u *Unit) MaxStamina() int    { return u.Stats.maxStamina }
func (u *Unit) Experience() int    { return u.Stats.exp }
func (u *Unit) MaxExperience() int { return u.Stats.maxExp }

func (u *Unit) updateStats() {
	if u.Stats.health < 0 {
		u.Stats.health = 0
	}
	if u.Stats.stamina < 0 {
		u.Stats.stamina = 0
	}
	if u.action.animation.tick%u.Stats.healthSpeedRecover == 0 && u.Stats.health < u.Stats.maxHealth {
		u.Stats.health++
	}
	if u.action.animation.tick%u.Stats.staminaSpeedRecover == 0 && u.Stats.stamina < u.Stats.maxStamina {
		u.Stats.stamina++
	}
}
