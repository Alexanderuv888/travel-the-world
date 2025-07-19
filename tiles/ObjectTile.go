package tiles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ObjectTile struct {
	Sx, Sy float64
	Tx, Ty float64
	W, H   int
	Img    *ebiten.Image
}

func NewObjectTile(tx, ty float64, w, h int, img *ebiten.Image) *ObjectTile {
	return &ObjectTile{
		Sx:  1,
		Sy:  1,
		Tx:  tx,
		Ty:  ty,
		W:   w,
		H:   h,
		Img: img}
}

func (o *ObjectTile) Width() float64 {
	return float64(o.W)
}

func (o *ObjectTile) Height() float64 {
	return float64(o.H)
}

func (o *ObjectTile) ScreenY() float64 {
	return o.Ty
}

func (o *ObjectTile) Draw(screen *ebiten.Image, cameraX, cameraY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(o.Sx, o.Sy)
	op.GeoM.Translate(o.Tx-cameraX, o.Ty-cameraY)
	screen.DrawImage(o.Img, op)
}
