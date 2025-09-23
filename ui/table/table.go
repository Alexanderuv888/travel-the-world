package table

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Table struct {
	X, Y, Width, Height int
	Headers             []string
	Rows                []*Row
	Style               Style
	selectedRowIndex    int
	visible             bool // Добавлено поле для видимости таблицы
	scroll              image.Rectangle
	minVisibleIndex     int
	maxVisibleIndex     int
	maxVisibleRows      int
	cursorY             int
	deltaY              int
	scrollIsDragging    bool
}

func NewTable(X, Y, Width, Height int) *Table {
	scroll := image.Rectangle{Min: image.Point{X: X + Width + 10, Y: Y + 5}, Max: image.Point{X: X + Width + 20, Y: Y + 35}}
	return &Table{
		X:                X,
		Y:                Y,
		Width:            Width,
		Height:           Height,
		scroll:           scroll,
		minVisibleIndex:  0,
		maxVisibleIndex:  7,
		maxVisibleRows:   7, // Максимальное количество видимых строк
		Style:            DefaultStyle(),
		selectedRowIndex: -1, // Изначально ни одна строка не выбрана
	}
}

func (t *Table) SetVisible(visible bool) {
	t.visible = visible
}

func (t *Table) SetStyle(style Style) {
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

func (t *Table) SelectedRowIndex() int {
	return t.selectedRowIndex
}

func (t *Table) ResetRowIndex() {
	if t.selectedRowIndex >= 0 && t.selectedRowIndex < len(t.Rows) {
		t.Rows[t.selectedRowIndex].Selected = false
	}
	t.selectedRowIndex = -1
}

func (t *Table) Update() {
	if !t.visible {
		return
	}
	t.UpdateScroll()
	for i := t.minVisibleIndex; i < t.maxVisibleIndex; i++ {
		if t.Rows[i].HoveredWith(ebiten.CursorPosition()) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			t.ResetRowIndex()
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

	x := 5 + t.X
	y := 5 + t.Y

	for i := t.minVisibleIndex; i < t.maxVisibleIndex; i++ {
		t.Rows[i].X = x
		t.Rows[i].Y = y
		t.Rows[i].Draw(dst)
		y += t.Rows[i].Height
	}
	t.drawScroll(dst)
}

func (t *Table) UpdateScroll() {
	if t.maxVisibleIndex > len(t.Rows) || t.minVisibleIndex < 0 {
		t.minVisibleIndex = 0
		t.maxVisibleIndex = len(t.Rows)
		return
	}
	if t.hoverScroll(ebiten.CursorPosition()) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			_, cursorY := ebiten.CursorPosition()
			t.cursorY = cursorY
			t.scrollIsDragging = true
		}
	}

	if t.scrollIsDragging {
		_, cursorY := ebiten.CursorPosition()
		t.deltaY = cursorY - t.cursorY
		t.cursorY = cursorY
		t.scroll.Min.Y += t.deltaY
		t.scroll.Max.Y += t.deltaY

		rowHeight := t.Rows[0].Height

		minY := t.Y
		maxY := t.Y + rowHeight*t.maxVisibleRows
		if t.scroll.Min.Y < minY {
			t.scroll.Min.Y = minY
			t.scroll.Max.Y = t.scroll.Min.Y + 30
		}
		if t.scroll.Max.Y > maxY {
			t.scroll.Max.Y = maxY
			t.scroll.Min.Y = t.scroll.Max.Y - 30
		}

		delta := maxY / len(t.Rows)

		d := (t.scroll.Min.Y + t.scroll.Dx()/2) / delta
		t.minVisibleIndex = d
		t.maxVisibleIndex = d + t.maxVisibleRows
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		t.scrollIsDragging = false
	}

	if t.minVisibleIndex < 0 {
		t.minVisibleIndex = 0
	}
	if t.maxVisibleIndex > len(t.Rows) {
		t.maxVisibleIndex = len(t.Rows)
	}
	if len(t.Rows) > t.maxVisibleRows {
		if t.minVisibleIndex > len(t.Rows)-t.maxVisibleRows {
			t.minVisibleIndex = len(t.Rows) - t.maxVisibleRows
		}
		if t.maxVisibleIndex < t.minVisibleIndex+t.maxVisibleRows {
			t.maxVisibleIndex = t.minVisibleIndex + t.maxVisibleRows
		}
	}
}

func (t *Table) hoverScroll(x, y int) bool {
	mp := image.Point{X: x, Y: y}
	return mp.In(t.scroll)
}

func (t *Table) drawScroll(dst *ebiten.Image) {
	if t.maxVisibleRows > len(t.Rows) {
		return
	}
	scroll := ebiten.NewImage(t.scroll.Dx(), t.scroll.Dy())
	scroll.Fill(color.RGBA{100, 100, 100, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.scroll.Min.X), float64(t.scroll.Min.Y))
	dst.DrawImage(scroll, op)
}
