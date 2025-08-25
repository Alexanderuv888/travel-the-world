package ui

import (
	"image/color"
	"log"
	"os"
	"path/filepath"
	text "travel-the-world/ui/text"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SaveMenu struct {
	X, Y, Width, Height int
	Visible             bool
	files               []SaveFile
	selectedIndex       int
	hoverIndex          int

	btnOverwrite *Button
	btnNewSave   *Button
	btnDelete    *Button

	ShowNewSaveDialog bool
	newSaveName       string
	btnOk             *Button
	btnCancel         *Button
}

func (m *SaveMenu) initCallBackFunckions() {
	onNewSave = new(func())
	*onNewSave = func() {
		m.ShowNewSaveDialog = true
	}
	onOverwrite = new(func())
	*onOverwrite = func() {
		if m.selectedIndex >= 0 && SaveGameHandler != nil {
			SaveGameHandler(m.files[m.selectedIndex].Name)
		}
	}
	onDelete = new(func())
	*onDelete = func() {
		if m.selectedIndex < 0 {
			return
		}
		// формируем путь с расширением (предполагаем .json)
		name := m.files[m.selectedIndex].Name
		path := filepath.Join("save", "saves", name+filepath.Ext(name+".json"))
		// Если имя уже содержало расширение (маловероятно, т.к. вы храните Name без расширения) — поправка:
		if filepath.Ext(m.files[m.selectedIndex].Name) != "" {
			path = filepath.Join("save", "saves", m.files[m.selectedIndex].Name)
		}
		if err := os.Remove(path); err != nil {
			log.Printf("failed to remove save file %s: %v", path, err)
			return
		}
		// Обновляем список файлов и сбрасываем индексы
		m.files = LoadSaveFiles()
		m.selectedIndex = -1
		m.hoverIndex = -1
	}
	onOkBtn = new(func())
	*onOkBtn = func() {
		if m.newSaveName != "" {
			if SaveGameHandler != nil {
				SaveGameHandler(m.newSaveName)
			}
			m.ShowNewSaveDialog = false
			m.newSaveName = ""
		}
	}
	onCancelBtn = new(func())
	*onCancelBtn = func() {
		m.ShowNewSaveDialog = false
		m.newSaveName = ""
	}
}

func NewSaveMenu(X, Y, Width, Height int) *SaveMenu {
	m := &SaveMenu{
		selectedIndex: -1,
		hoverIndex:    -1,
		X:             X,
		Y:             Y,
		Width:         Width,
		Height:        Height,
	}

	// Загружаем список файлов сохранений
	m.files = LoadSaveFiles()

	// Кнопки
	m.btnNewSave = NewButton(m.X+m.Width-40-450, m.Y+m.Height+10, 150, 40, "New save", onNewSave)

	m.btnOverwrite = NewButton(m.X+m.Width-20-300, m.Y+m.Height+10, 150, 40, "Overwrite", onOverwrite)
	m.btnDelete = NewButton(m.X+m.Width-150, m.Y+m.Height+10, 150, 40, "Delete", onDelete)

	// Кнопки диалога
	m.btnOk = NewButton(m.X+m.Width/2-150+5, m.Y+m.Height/2-60+80, 80, 30, "OK", onOkBtn)
	m.btnCancel = NewButton(m.X+m.Width/2-50+5, m.Y+m.Height/2-60+80, 80, 30, "Cancel", onCancelBtn)

	return m
}

func (m *SaveMenu) LoadFiles() {
	m.files = LoadSaveFiles()
}

func (m *SaveMenu) Update() {
	if !m.Visible {
		return
	}
	if m.ShowNewSaveDialog {
		m.btnOk.Update()
		m.btnCancel.Update()

		// Обновляем строку ввода имени файла
		for _, r := range ebiten.AppendInputChars(nil) {
			if r == '\n' || r == '\r' {
				continue
			}
			m.newSaveName += string(r)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(m.newSaveName) > 0 {
			m.newSaveName = m.newSaveName[:len(m.newSaveName)-1]
		}

		return
	}

	// Определяем наведение
	cx, cy := ebiten.CursorPosition()
	m.hoverIndex = -1
	for i := range m.files {
		rowY := m.Y + 5 + i*32
		if cx >= m.X+5 && cx <= m.X+m.Width+5 && cy >= rowY && cy <= rowY+32 {
			m.hoverIndex = i
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				m.selectedIndex = i
			}
		}
	}

	// Обновляем кнопки
	m.btnNewSave.Update()
	m.btnOverwrite.Update()
	m.btnDelete.Update()
}

func (m *SaveMenu) Draw(screen *ebiten.Image) {
	if !m.Visible {
		return
	}

	clr := color.RGBA{50, 50, 50, 255}

	dst := ebiten.NewImage(int(m.Width), int(m.Height))
	dst.Fill(clr)

	dstOp := &ebiten.DrawImageOptions{}
	dstOp.GeoM.Translate(float64(m.X), float64(m.Y))

	// Таблица
	for i, file := range m.files {
		rowY := float64(5 + i*32)

		clr := color.RGBA{50, 50, 50, 255} // обычная строка
		if i == m.hoverIndex {
			clr = color.RGBA{60, 60, 60, 255} // hover
		}
		if i == m.selectedIndex {
			clr = color.RGBA{80, 80, 80, 255} // selected
		}

		sub := ebiten.NewImage(dst.Bounds().Dx()-10, 30)
		sub.Fill(clr)

		text.DrawText(sub, file.Name, color.White, 20, 5, text.Left)
		text.DrawText(sub, file.ModTime.Format("2006-01-02 15:04"), color.White, 20, 5, text.Right)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(5, rowY)

		dst.DrawImage(sub, op)
	}

	// Кнопки
	m.btnNewSave.Draw(screen)
	m.btnOverwrite.Draw(screen)
	m.btnDelete.Draw(screen)

	if m.ShowNewSaveDialog {
		dstOp.ColorScale.ScaleAlpha(float32(0.2))
	}
	screen.DrawImage(dst, dstOp)

	// Диалог нового сохранения
	if m.ShowNewSaveDialog {

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

		m.btnOk.Draw(screen)
		m.btnCancel.Draw(screen)
	}

}
