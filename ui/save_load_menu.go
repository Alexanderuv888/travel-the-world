package ui

import (
	"image/color"
	"travel-the-world/ui/table"

	text "travel-the-world/ui/text"

	"github.com/hajimehoshi/ebiten/v2"
)

type SaveLoadMenu struct {
	X, Y, Width, Height int
	table               *table.Table
	buttons             []*Button
	files               []SaveFile
	selectedIndex       int
	hoverIndex          int
	visible             bool
}

func NewSaveLoadMenu(X, Y, Width, Height int) *SaveLoadMenu {
	m := &SaveLoadMenu{
		selectedIndex: -1,
		hoverIndex:    -1,
		X:             X,
		Y:             Y,
		Width:         Width,
		Height:        Height,
	}
	m.table = &table.Table{
		X:      X + 5,
		Y:      Y + 5,
		Width:  Width - 10,
		Height: Height - 200,
	}
	m.table.SetStyle(table.DefaultStyle())
	m.files = LoadSaveFiles()

	return m
}

func (m *SaveLoadMenu) SetVisible(val bool) {
	m.visible = val
	m.table.SetVisible(val)
}

func (m *SaveLoadMenu) SetFiles(files []SaveFile) {
	m.files = files
	m.table.Rows = nil
	for _, file := range m.files {
		name := file.Name
		date := file.ModTime.Format("2006-01-02 15:04")
		dateWidth := m.table.Width / 4
		nameWidth := m.table.Width - dateWidth
		height := m.table.Style.GetTextSize() + 10
		style := table.DefaultStyle()
		style.TextAlign = text.Right
		nameCell := table.NewCell(nameWidth, height, name, m.table.Style)
		dateCell := table.NewCell(dateWidth, height, date, style)

		m.table.AddRow(table.NewRow(nameCell, dateCell))
	}
}

func (m *SaveLoadMenu) SetButtons(buttons ...*Button) {
	m.buttons = buttons
}

func (m *SaveLoadMenu) Update() {
	if !m.visible {
		return
	}
	m.table.Update()
	for _, btn := range m.buttons {
		btn.Update()
	}
}

func (m *SaveLoadMenu) Draw(screen *ebiten.Image) {
	if !m.visible {
		return
	}

	clr := color.RGBA{50, 50, 50, 255}

	dst := ebiten.NewImage(m.Width, m.Height)
	dst.Fill(clr)

	dstOp := &ebiten.DrawImageOptions{}
	dstOp.GeoM.Translate(float64(m.X), float64(m.Y))
	m.table.Draw(dst)
	for _, btn := range m.buttons {
		btn.Draw(dst)
	}

	screen.DrawImage(dst, dstOp)
}
