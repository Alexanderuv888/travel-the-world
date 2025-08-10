package tiles

import (
	"image"
	"travel-the-world/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	ScreenY() float64                                                                    // для сортировки по глубине
	Draw(screen *ebiten.Image, camera *camera.Camera, playerPos image.Point, radius int) // как отрисовывать себя
}
