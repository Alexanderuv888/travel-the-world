package table

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Row struct {
	X, Y, Width, Height int
	d                   image.Rectangle
	Cells               []*Cell
	Style               *Style
	Hovered             bool
	Selected            bool
}

func NewRow(cells ...*Cell) *Row {
	r := Row{}
	r.Cells = cells
	r.Width = 0
	r.Height = 0
	for _, cell := range cells {
		r.Width += cell.Width
		if r.Height < cell.Height {
			r.Height = cell.Height
		}
	}
	return &r
}

func (r *Row) Update() {
	for _, cell := range r.Cells {
		cell.Hovered = r.Hovered
		cell.Selected = r.Selected
		cell.Update()
	}
}

func (r *Row) Draw(dst *ebiten.Image) {
	x := r.X
	for _, cell := range r.Cells {
		cell.X = x
		cell.Y = r.Y
		cell.Draw(dst)
		x += cell.Width
	}
}

func (r *Row) HoveredWith(x, y int) bool {
	mp := image.Point{X: x, Y: y}
	r.d = image.Rectangle{image.Point{r.X, r.Y}, image.Point{r.X + r.Width, r.Y + r.Height}}
	r.Hovered = mp.In(r.d)
	return r.Hovered
}
