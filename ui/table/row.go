package table

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Row struct {
	d        image.Rectangle
	Cells    []Cell
	Style    Style
	hovered  bool
	selected bool
}

func createRowFromStrings(contents []string) Row {
	cells := make([]Cell, len(contents))
	for i, content := range contents {
		cells[i] = Cell{
			Content: content,
		}
	}
	return Row{
		Cells: cells,
	}
}

func (r *Row) Update() {
	r.hovered = false
	r.selected = false
	for _, cell := range r.Cells {
		cell.Update()
		if cell.Hovered {
			r.hovered = true
		}
		if cell.Selected {
			r.selected = true
		}
	}
}

func (r *Row) Draw(dst *ebiten.Image) {
	for _, cell := range r.Cells {
		cell.Draw(dst)
	}
}
