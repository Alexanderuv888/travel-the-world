package unit

import (
	"image"
	"travel-the-world/assets"
	"travel-the-world/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

func (u *Unit) Update(screen *ebiten.Image, dq *tiles.DrawQueue, levelDimentions image.Point, am *assets.Manager) {
	u.updateAnimation(screen, levelDimentions, am)
	dq.Add(u.animation.CTile)
}
