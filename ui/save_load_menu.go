package ui

import (
	"travel-the-world/ui/table"

	text "travel-the-world/ui/text"

	"github.com/hajimehoshi/ebiten/v2"
)

type SaveLoadMenu struct {
	X, Y, Width, Height int
	table               *table.Table
	buttonsBar          *ButtonsBar
	visible             bool
	showNewSaveDialog   bool
	newSaveName         string
	newSaveDialog       *NewSaveDialog
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

	m.newSaveDialog = NewNewSaveDialog(X+Width/2-250, Y+Height/2-60, 500, 200)
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

func (m *SaveLoadMenu) SaveFile() {
	if m.newSaveDialog.newSaveName != "" {
		if SaveGameHandler != nil {
			SaveGameHandler(m.newSaveDialog.newSaveName)
		}
		m.SetVisibleNewSaveDilog(false)
		m.SetFiles(LoadSaveFiles())
		m.resetIndex()
	}
}

func (m *SaveLoadMenu) SetVisibleNewSaveDilog(visible bool) {
	m.showNewSaveDialog = visible
	m.newSaveDialog.newSaveName = ""
	m.newSaveDialog.SetVisible(visible)
}

func (m *SaveLoadMenu) SetButtons(buttons []*Button) {
	m.buttonsBar.SetButtons(buttons)
}

func (m *SaveLoadMenu) Update() {
	if !m.visible {
		return
	}
	if m.showNewSaveDialog {
		m.newSaveDialog.Update()
	} else {
		m.table.Update()
		m.buttonsBar.Update()
	}
}

func (m *SaveLoadMenu) Draw(screen *ebiten.Image) {
	if !m.visible {
		return
	}
	m.table.Draw(screen)
	m.buttonsBar.Draw(screen)
	if m.showNewSaveDialog {
		m.newSaveDialog.Draw(screen)
	}
}
