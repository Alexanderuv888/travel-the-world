package ui

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	//log.Println("ui/buttons_initializatoin.go init()")
}

var (
	onNewSave   *func()
	onOverwrite *func()
	onLoad      *func()
	onDelete    *func()
	onClose     *func()
	onOkBtn     *func()
	onCancelBtn *func()
)

var (
	btnOverwrite *Button
	btnNewSave   *Button
	btnLoad      *Button
	btnDelete    *Button
	btnClose     *Button
	btnOk        *Button
	btnCancel    *Button
)

var SaveMenuButtons []*Button
var LoadMenuButtons []*Button
var NewSaveDilogButtons []*Button

func initButtons() {

	onNewSave = new(func())
	onOverwrite = new(func())
	onLoad = new(func())
	onDelete = new(func())
	onClose = new(func())
	onOkBtn = new(func())
	onCancelBtn = new(func())

	btnNewSave = DefButton("New save", onNewSave)
	btnOverwrite = DefButton("Overwrite", onOverwrite)
	btnLoad = DefButton("Load", onLoad)
	btnDelete = DefButton("Delete", onDelete)
	btnClose = DefButton("Close", onClose)
	btnOk = DefButton("OK", onOkBtn)
	btnCancel = DefButton("Cancel", onCancelBtn)

	SaveMenuButtons = []*Button{btnNewSave, btnOverwrite, btnDelete, btnClose}
	LoadMenuButtons = []*Button{btnLoad, btnDelete, btnClose}
	NewSaveDilogButtons = []*Button{btnOk, btnCancel}
}

// обработчик сохранения - его нужно зарегистрировать снаружи (game.SaveProgress)
var SaveGameHandler func(filename string)

var LoadGameHandler func(filename string)

// Экспортируем регистратор обработчика сохранения
func SetSaveGameHandler(cb func(filename string)) {
	SaveGameHandler = cb
}

func SetLoadHandler(cb func(filename string)) {
	LoadGameHandler = cb
}

func (m *SaveLoadMenu) InitSaveLoadCallBackFunckions() {
	*onNewSave = func() {
		m.ShowNewSaveDialog = true
	}
	*onOverwrite = func() {
		if m.selectedIndex() >= 0 && SaveGameHandler != nil {
			SaveGameHandler(m.selectedFileName())
		}
	}
	*onLoad = func() {
		if m.selectedIndex() >= 0 && LoadGameHandler != nil {
			LoadGameHandler(m.selectedFileName())
			m.visible = false
			m.table.SetVisible(false)
		}
	}
	*onDelete = func() {
		if m.selectedIndex() < 0 {
			return
		}
		// формируем путь с расширением (предполагаем .json)
		name := m.selectedFileName()
		path := filepath.Join("save", "saves", name+filepath.Ext(name+".json"))
		// Если имя уже содержало расширение (маловероятно, т.к. вы храните Name без расширения) — поправка:
		if filepath.Ext(name) != "" {
			path = filepath.Join("save", "saves", name)
		}
		if err := os.Remove(path); err != nil {
			log.Printf("Failed to remove save file %s: %v", path, err)
			return
		}
		// Обновляем список файлов и сбрасываем индексы
		m.SetFiles(LoadSaveFiles())
		m.resetIndex()
	}
	*onClose = func() {
		m.visible = false
		m.SetVisible(false)
		m.ShowNewSaveDialog = false
		m.newSaveName = ""
		m.resetIndex()
	}
}

func (m *SaveMenu) initConfirmationDialogCallbacks() {
	*onOkBtn = func() {
		if m.newSaveName != "" {
			if SaveGameHandler != nil {
				SaveGameHandler(m.newSaveName)
			}
			m.ShowNewSaveDialog = false
			m.newSaveName = ""
		}
	}
	*onCancelBtn = func() {
		m.ShowNewSaveDialog = false
		m.newSaveName = ""
	}
}
