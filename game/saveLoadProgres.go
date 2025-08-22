package game

import (
	"log"
	"path/filepath"
	"travel-the-world/save"
)

func (g *Game) SaveProgress(fileName string) {
	data := save.SaveData{
		Health:              g.Player.Health(),
		MaxHealth:           g.Player.MaxHealth(),
		HealthSpeedRecover:  g.Player.HealthSpeedRecover(),
		Stamina:             g.Player.Stamina(),
		MaxStamina:          g.Player.MaxStamina(),
		StaminaSpeedRecover: g.Player.StaminaSpeedRecover(),
		Damage:              g.Player.DamagePoint(),
		AttackDistance:      g.Player.AttackDistance(),
		Vision:              g.Player.Vision(),
		Speed:               g.Player.Speed(),
		Level:               g.Player.Level(),
		MaxExp:              g.Player.MaxExperience(),
		Exp:                 g.Player.Experience(),
	}
	file := filepath.Join(fileName + filepath.Ext(fileName+".json"))
	if err := save.SaveGame(data, file); err != nil {
		log.Fatal("Ошибка сохранения прогресса", err)
	}
}

func (g *Game) LoadProgress() {
	data, err := save.LoadGame("/Users/vols/Projects/travel-the-world/save.json")
	if err != nil {
		log.Fatal("Ошибка загрузки прогресса", err)
		return
	}
	g.Player.LoadStats(data)
}
