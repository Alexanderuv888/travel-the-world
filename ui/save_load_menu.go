package ui

import (
	"image/color"
	"travel-the-world/ui/table"

	text "travel-the-world/ui/text"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SaveLoadMenu struct {
	X, Y, Width, Height int
	table               *table.Table
	buttonsBar          *ButtonsBar
	visible             bool
	ShowNewSaveDialog   bool
	newSaveName         string
}

func NewSaveLoadMenu(X, Y, Width, Height int) *SaveLoadMenu {
	m := &SaveLoadMenu{
		X:      X,
		Y:      Y,
		Width:  Width,
		Height: Height,
	}
	m.table = table.NewTable(X+5, Y+5, Width-10, Height-200)
	m.buttonsBar = NewButtonsBar(X+5, Y+Height-200, Width-10, 200)
	m.buttonsBar.SetOrientation(OrientHorizontal)
	m.buttonsBar.SetAlignment(AlignRight)
	m.buttonsBar.SetPadding(10)
	m.buttonsBar.SetSpacing(10)
	return m
}

func (m *SaveLoadMenu) SetVisible(val bool) {
	m.visible = val
	m.table.SetVisible(val)
	m.buttonsBar.SetVisible(val)
}

func (m *SaveLoadMenu) SetFiles(files []SaveFile) {
	m.table.Rows = nil
	for _, file := range files {
		name := file.Name
		date := file.ModTime.Format("2006-01-02 15:04")
		dateWidth := m.table.Width / 4
		nameWidth := m.table.Width - dateWidth
		height := m.table.Style.GetTextSize() + 15
		nameCell := table.NewCell(nameWidth, height, name, m.table.Style)
		dateCell := table.NewCell(dateWidth, height, date, m.table.Style)
		dateCell.SetAlign(text.Right)
		m.table.AddRow(table.NewRow(nameCell, dateCell))
	}
	m.table.UpdateScroll()
}

func (m *SaveLoadMenu) selectedFileName() string {
	if m.selectedIndex() >= 0 && m.selectedIndex() < len(m.table.Rows) {
		return m.table.Rows[m.selectedIndex()].Cells[0].Content
	}
	return ""
}

func (m *SaveLoadMenu) selectedIndex() int {
	return m.table.SelectedRowIndex()
}

func (m *SaveLoadMenu) resetIndex() {
	m.table.ResetRowIndex()
}

func (m *SaveLoadMenu) SetButtons(buttons []*Button) {
	m.buttonsBar.SetButtons(buttons)
}

func (m *SaveLoadMenu) Update() {
	if !m.visible {
		return
	}
	m.table.Update()
	m.buttonsBar.Update()
}

func (m *SaveLoadMenu) Draw(screen *ebiten.Image) {
	if !m.visible {
		return
	}
	m.table.Draw(screen)
	m.buttonsBar.Draw(screen)
	if m.ShowNewSaveDialog {
		clr := color.RGBA{50, 50, 50, 255} // обычная строка
		clr1 := color.RGBA{30, 30, 30, 240}
		vector.DrawFilledRect(screen, float32(m.X+m.Width/2-150), float32(m.Y+m.Height/2-60), 300, 120, clr1, false)

		row1 := ebiten.NewImage(290, 12)
		row1.Fill(clr1)
		op1 := &ebiten.DrawImageOptions{}
		op1.GeoM.Translate(float64(m.X+m.Width/2-150+5), float64(m.Y+m.Height/2-60+20))
		text.DrawText(row1, "Enter save name:", color.White, 10, 5, text.Right)
		screen.DrawImage(row1, op1)

		row2 := ebiten.NewImage(290, 32)
		row2.Fill(clr)
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(m.X+m.Width/2-150+5), float64(m.Y+m.Height/2-60+40))
		text.DrawText(row2, m.newSaveName, color.RGBA{200, 200, 200, 255}, 20, 5, text.Right)
		screen.DrawImage(row2, op2)

		//m.btnOk.Draw(screen)
		//m.btnCancel.Draw(screen)
	}
}
