package table

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Table struct {
	X, Y, Width, Height int
	Headers             []string
	Rows                []*Row
	D                   image.Rectangle
	Style               *Style
	selectedRowIndex    int
	visible             bool // Добавлено поле для видимости таблицы
}

func (t *Table) SetVisible(visible bool) {
	t.visible = visible
}

func (t *Table) SetStyle(style *Style) {
	t.Style = style
	for i := range t.Rows {
		t.Rows[i].Style = style
		for j := range t.Rows[i].Cells {
			t.Rows[i].Cells[j].style = style
		}
	}
}

func (t *Table) AddRow(row *Row) {
	row.Style = t.Style
	t.Rows = append(t.Rows, row)
}

func (t *Table) SetHeaders(headers []string) {
	t.Headers = headers
}

func (t *Table) Update() {
	if !t.visible {
		return
	}

	cx, cy := ebiten.CursorPosition()
	x := cx - t.X
	y := cy - t.Y
	for i := range t.Rows {
		if t.Rows[i].HoveredWith(x, y) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			t.Rows[t.selectedRowIndex].Selected = false
			t.Rows[i].Selected = true
			t.selectedRowIndex = i
		}
		t.Rows[i].Update()
	}
}

func (t *Table) Draw(dst *ebiten.Image) {
	if !t.visible {
		return
	}

	x := 5
	y := 5

	// Draw rows
	for _, row := range t.Rows {
		row.X = x
		row.Y = y
		row.Draw(dst)
		y += row.Height
	}
}
