package unit

type Stats struct {
	health           int
	maxHealth        int
	damage           int
	AttackDistance   float64
	Vision           float64
	speed            float64
	level            int
	expTillNextLevel int
	expireance       int
}

func NewStats(health, maxHealth, damage int, attackDisance, vision, speed float64) *Stats {
	return &Stats{
		health:           health,
		maxHealth:        maxHealth,
		damage:           damage,
		AttackDistance:   attackDisance,
		Vision:           vision,
		speed:            speed,
		level:            1,
		expTillNextLevel: 10,
		expireance:       0}
}

func (u *Unit) AddExp(exp int) {
	s := u.Stats
	u.Stats.expireance += exp
	if u.Stats.expireance >= u.Stats.expTillNextLevel {
		expLeft := u.Stats.expireance - u.Stats.expTillNextLevel
		u.levelUp()
		u.Stats = s
		u.AddExp(expLeft)
	}
}

func (u *Unit) levelUp() {
	u.Stats.level++
	u.Stats.expireance = 0
	u.Stats.expTillNextLevel = u.Stats.expTillNextLevel + factorial(u.Stats.level)
	u.Stats.maxHealth += 2
	u.Stats.health = u.Stats.maxHealth
	u.Stats.damage += 1
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
