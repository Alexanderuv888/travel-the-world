package tiles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	Width() float64
	Height() float64
	ScreenY() float64                                    // для сортировки по глубине
	Draw(screen *ebiten.Image, cameraX, cameraY float64) // как отрисовывать себя
}
