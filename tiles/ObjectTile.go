package tiles

import (
	"image"
	"math"
	"travel-the-world/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

type ObjectTile struct {
	Sx, Sy float64
	Tx, Ty float64
	Img    *ebiten.Image
}

func NewObjectTile(tx, ty float64, img *ebiten.Image) *ObjectTile {
	return &ObjectTile{
		Sx:  1,
		Sy:  1,
		Tx:  tx,
		Ty:  ty,
		Img: img}
}

func (o *ObjectTile) Centr() image.Point {
	return image.Point{
		X: int(o.Tx) + o.Img.Bounds().Dx()/2,
		Y: int(o.Ty) + o.Img.Bounds().Dy()/2,
	}
}

func (o *ObjectTile) ScreenY() float64 {
	return o.Ty
}

func (o *ObjectTile) Draw(screen *ebiten.Image, camera *camera.Camera, playerPos image.Point, radius int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.Tx, o.Ty)
	camera.Apply(op)
	alpha := 1.0
	if isOverlapping2(playerPos, o.Centr(), radius) {
		alpha = 0.2
	}
	op.ColorScale.ScaleAlpha(float32(alpha))
	screen.DrawImage(o.Img, op)
}

func isOverlapping(p1, p2 image.Point, radius int) bool {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return dx*dx+dy*dy < radius*radius
}

func isOverlapping2(p1, p2 image.Point, radius int) bool {
	dx := float64(p2.X - p1.X)
	dy := float64(p2.Y - p1.Y)
	return math.Sqrt(dx*dx+dy*dy) < float64(radius)
}
