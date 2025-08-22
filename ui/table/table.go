package table

import (
	"image"
	"image/color"
	"travel-the-world/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Table struct {
	Headers []string
	Rows    []Row
	d       image.Rectangle
	style   Style
	visible bool // Добавлено поле для видимости таблицы
}

func (t *Table) SetVisible(visible bool) {
	t.visible = visible
}

func (t *Table) SetStyle(style Style) {
	t.style = style
	for i := range t.Rows {
		t.Rows[i].Style = style
		for j := range t.Rows[i].Cells {
			t.Rows[i].Cells[j].style = style
		}
	}
}

func (t *Table) AddRow(row Row) {
	t.Rows = append(t.Rows, row)
}
func (t *Table) SetHeaders(headers []string) {
	t.Headers = headers
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Update() {
	if !t.visible {
		return
	}
	for i := range t.Rows {
		t.Rows[i].Update()
	}
}

func (t *Table) Draw(dst *ebiten.Image) {
	if !t.visible {
		return
	}

	x := dst.Bounds().Dx() / 4
	y := dst.Bounds().Dy() / 4

	width := dst.Bounds().Dx() / 2
	height := dst.Bounds().Dy() / 2

	t.d = image.Rectangle{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: x + width, Y: y + height},
	}

	headerHeight := 30
	//rowHeight := 25

	// Draw headers
	for i, header := range t.Headers {
		// Draw header background
		headerRect := ebiten.NewImage(t.d.Dx()/len(t.Headers), headerHeight)
		headerRect.Fill(color.RGBA{200, 200, 200, 255})
		ui.DrawText(headerRect, header, color.Black, 10, 5, ui.Center)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(t.d.Min.X+i*(t.d.Dx()/len(t.Headers))), float64(t.d.Min.Y))
		dst.DrawImage(headerRect, op)
	}

	// Draw rows
	for _, row := range t.Rows {
		for _, cell := range row.Cells {
			cell.Draw(dst)
		}
	}
}
