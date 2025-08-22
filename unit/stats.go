package unit

import (
	"encoding/json"
	"log"
	"travel-the-world/save"
)

type Stats struct {
	health              int     `json:"health"`
	maxHealth           int     `json:"maxHealth"`
	healthSpeedRecover  int     `json:"healthSpeedRecover"`
	stamina             int     `json:"stamina"`
	maxStamina          int     `json:"maxStamina"`
	staminaSpeedRecover int     `json:"staminaSpeedRecover"`
	damage              int     `json:"damage"`
	AttackDistance      float64 `json:"AttackDistance"`
	Vision              float64 `json:"Vision"`
	speed               float64 `json:"speed"`
	level               int     `json:"level"`
	maxExp              int     `json:"maxExp"`
	exp                 int     `json:"exp"`
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

func (u *Unit) Health() int              { return u.Stats.health }
func (u *Unit) MaxHealth() int           { return u.Stats.maxHealth }
func (u *Unit) HealthSpeedRecover() int  { return u.Stats.healthSpeedRecover }
func (u *Unit) Stamina() int             { return u.Stats.stamina }
func (u *Unit) MaxStamina() int          { return u.Stats.maxStamina }
func (u *Unit) StaminaSpeedRecover() int { return u.Stats.staminaSpeedRecover }
func (u *Unit) DamagePoint() int         { return u.Stats.damage }
func (u *Unit) AttackDistance() float64  { return u.Stats.AttackDistance }
func (u *Unit) Vision() float64          { return u.Stats.Vision }
func (u *Unit) Speed() float64           { return u.Stats.speed }
func (u *Unit) Level() int               { return u.Stats.level }
func (u *Unit) Experience() int          { return u.Stats.exp }
func (u *Unit) MaxExperience() int       { return u.Stats.maxExp }

func (u *Unit) LoadStats(data save.SaveData) {
	u.Stats = &Stats{
		health:              data.Health,
		maxHealth:           data.MaxHealth,
		healthSpeedRecover:  data.HealthSpeedRecover,
		stamina:             data.Stamina,
		maxStamina:          data.MaxStamina,
		staminaSpeedRecover: data.StaminaSpeedRecover,
		damage:              data.Damage,
		AttackDistance:      data.AttackDistance,
		Vision:              data.Vision,
		speed:               data.Speed,
		level:               data.Level,
		maxExp:              data.MaxExp,
		exp:                 data.Exp}
}

func (u *Unit) helthLeftInPercent() float64 {
	return float64(u.Health()) / float64(u.Stats.maxHealth)
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

func (u *Unit) updateStats() {
	if u.Stats.health < 0 {
		u.Stats.health = 0
	}
	if u.Stats.stamina < 0 {
		u.Stats.stamina = 0
	}
	if u.action.animation.tick%u.Stats.healthSpeedRecover == 0 && u.Health() < u.Stats.maxHealth {
		u.Stats.health++
	}
	if u.action.animation.tick%u.Stats.staminaSpeedRecover == 0 && u.Stats.stamina < u.Stats.maxStamina {
		u.Stats.stamina++
	}
}

func (s *Stats) ConvertToJson() []byte {
	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Fatal("Ошибка при конвертации в json структуру:", err)
		return nil
	}
	return jsonData
}

func (s *Stats) ConvertFromJson(jsonData []byte) {
	var stats Stats
	err := json.Unmarshal(jsonData, stats)
	if err != nil {
		log.Fatal("Ошибка при конвертации из json структуры:", err)
		return
	}
	s = &stats
}
