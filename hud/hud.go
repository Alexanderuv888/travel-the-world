package hud

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

type HUD struct {
	Hero interface {
		Health() int
		MaxHealth() int
		Stamina() int
		MaxStamina() int
		Experience() int
		MaxExperience() int
	}
}

func NewHUD(hero any) *HUD {
	return &HUD{Hero: hero.(interface {
		Health() int
		MaxHealth() int
		Stamina() int
		MaxStamina() int
		Experience() int
		MaxExperience() int
	})}
}

func (h *HUD) Draw(screen *ebiten.Image) {
	const (
		barWidth  = 200
		barHeight = 12
		spacing   = 8
		margin    = 10
		borderRad = 4
	)
	w, _ := ebiten.WindowSize()

	x := float32(w - margin)
	y := float32(margin)

	drawBar := func(current, max int, y float32, col color.Color) {
		ratio := float32(current) / float32(max)
		filledWidth := float32(barWidth) * ratio

		// Фон
		vector.DrawFilledRect(screen, x, y, -float32(barWidth), float32(barHeight), color.RGBA{30, 30, 30, 220}, false)
		// Заполненная часть
		vector.DrawFilledRect(screen, x, y, -filledWidth, float32(barHeight), col, false)
	}

	drawBar(h.Hero.Health(), h.Hero.MaxHealth(), y, color.RGBA{200, 40, 40, 255})
	y += barHeight + spacing
	drawBar(h.Hero.Stamina(), h.Hero.MaxStamina(), y, color.RGBA{40, 200, 40, 255})
	y += barHeight + spacing
	drawBar(h.Hero.Experience(), h.Hero.MaxExperience(), y, color.RGBA{60, 120, 255, 255})
}
