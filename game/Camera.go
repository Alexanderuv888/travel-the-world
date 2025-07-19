package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	cameraSpeed = 3.0
)

type Camera struct {
	X, Y   float64
	vx, vy float64
	kx, ky float64
}

func NewCamera(x, y float64) *Camera {
	return &Camera{
		X: x,
		Y: y,
	}
}

func (camera *Camera) Update() {
	camera.vx, camera.vy = 0, 0
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		camera.vx = -cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		camera.vx = +cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		camera.vy = -cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		camera.vy = +cameraSpeed
	}

	xoff, yoff := ebiten.Wheel()

	camera.vx = -4 * xoff
	camera.vy = -4 * yoff

	camera.X += camera.vx
	camera.Y += camera.vy
}
