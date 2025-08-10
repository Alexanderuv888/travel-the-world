package tiles

import (
	"image"
	"travel-the-world/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

type CompositeTile struct {
	Sx, Sy float64
	Tx, Ty float64
	W, H   int
	Images []*ebiten.Image
}

func NewCompositeTile(tx, ty float64, w, h int, images []*ebiten.Image) *CompositeTile {
	return &CompositeTile{
		Sx:     1,
		Sy:     1,
		Tx:     tx,
		Ty:     ty,
		W:      w,
		H:      h,
		Images: images,
	}
}

func (o *CompositeTile) Width() float64 {
	return float64(o.W)
}

func (o *CompositeTile) Height() float64 {
	return float64(o.H)
}

func (o *CompositeTile) ScreenY() float64 {
	return o.Ty - float64(o.H)
}

func (o *CompositeTile) Draw(screen *ebiten.Image, camera *camera.Camera, playerPos image.Point, radius int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.Tx, o.Ty)
	camera.Apply(op)
	for _, img := range o.Images {
		screen.DrawImage(img, op)
	}
}
