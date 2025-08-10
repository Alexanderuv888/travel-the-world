package camera

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cameraSpeed = 3.0
)

type Camera struct {
	X, Y    float64
	vx, vy  float64
	kx, ky  float64
	Scale   float64
	ScaleTo float64
}

const (
	scaleMax   = 3
	scaleMin   = 1
	scaleSpeed = 0.01
)

func NewCamera(x, y float64) *Camera {
	return &Camera{
		X:     x,
		Y:     y,
		Scale: 1,
	}
}

func (camera *Camera) Pos() image.Point {
	return image.Point{int(camera.X), int(camera.Y)}
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

	camera.updateZoom()
}

func (c *Camera) updateZoom() {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		c.Scale += scaleSpeed
		if c.Scale > scaleMax {
			c.Scale = scaleMax
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		c.Scale -= scaleSpeed
		if c.Scale < scaleMin {
			c.Scale = scaleMin
		}
	}

}

func (c *Camera) Apply(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(-c.X, -c.Y)   // Сдвиг камеры
	op.GeoM.Scale(c.Scale, c.Scale) // Масштаб

	// Масштабирование относительно центра экрана
	//sw, sh := ebiten.WindowSize()
	//op.GeoM.Translate(float64(sw)/2, float64(sh)/2) // Центр
}

func (c *Camera) WorldToScreen(worldX, worldY float64) (screenX, screenY float64) {
	screenX = (worldX - c.X) * c.Scale
	screenY = (worldY - c.Y) * c.Scale
	return
}

// ScreenToWorld — экранные координаты → мировые
func (c *Camera) ScreenToWorld(screenX, screenY float64) (worldX, worldY float64) {
	worldX = screenX/c.Scale + c.X
	worldY = screenY/c.Scale + c.Y
	return
}

func (c *Camera) ScreenToWorldPoint(screenX, screenY int) image.Point {
	worldX := float64(screenX)/c.Scale + c.X
	worldY := float64(screenY)/c.Scale + c.Y
	return image.Point{int(worldX), int(worldY)}
}

func (c *Camera) ScreenToWorldRect(r image.Rectangle) image.Rectangle {
	pMin := c.ScreenToWorldPoint(r.Min.X, r.Min.Y)
	pMax := c.ScreenToWorldPoint(r.Max.X, r.Max.Y)
	return image.Rectangle{pMin, pMax}
}
