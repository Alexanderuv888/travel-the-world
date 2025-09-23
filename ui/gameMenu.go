package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	Buttons []*Button
	//Sfll    *SaveFilesListLayout
	//SavesTable *Table
	SaveMenu     SaveMenu
	SaveLoadMenu SaveLoadMenu
}

func NewMenu() *Menu {
	w, h := ebiten.WindowSize()
	wfloat := float64(w)
	hfloat := float64(h)
	onGameStart = new(func())
	onLoadGame = new(func())
	onSaveGame = new(func())
	onExit = new(func())

	startGameBtn := NewButton(100, 100, 200, 40, "New game", onGameStart)
	saveGameBtn := NewButton(100, 160, 200, 40, "Save game", onSaveGame)
	loadGameBtn := NewButton(100, 220, 200, 40, "Load game", onLoadGame)
	exitBtn := NewButton(100, 280, 200, 40, "Quit", onExit)

	saveMenu := NewSaveMenu(450, 100, int(wfloat/2), int(hfloat/3))
	saveLoadMenu := NewSaveLoadMenu(450, 100, int(wfloat/2), int(hfloat/3))
	saveLoadMenu.InitSaveLoadCallBackFunckions()
	saveLoadMenu.initConfirmationDialogCallbacks()
	m := &Menu{Buttons: []*Button{startGameBtn, loadGameBtn, saveGameBtn, exitBtn}, SaveMenu: *saveMenu, SaveLoadMenu: *saveLoadMenu}
	m.SaveMenu.initCallBackFunckions()
	return m
}

var (
	onGameStart *func()
	onLoadGame  *func()
	onSaveGame  *func()
	onExit      *func()
)

func (m *Menu) SetStartCallback(cb func()) {
	*onGameStart = cb
}

func (m *Menu) SetLoadGameCallback(cb func()) {
	*onLoadGame = cb
}

func (m *Menu) SetSaveGameCallback(cb func()) {
	*onSaveGame = cb
}

func (m *Menu) SetExitCallback(cb func()) {
	*onExit = cb
}

func (m *Menu) Update() {
	for _, b := range m.Buttons {
		b.isHovered(ebiten.CursorPosition())
		b.Update()
	}
	m.SaveMenu.Update()
	m.SaveLoadMenu.Update()
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for _, b := range m.Buttons {
		b.Draw(screen)
	}
	m.SaveMenu.Draw(screen)
	m.SaveLoadMenu.Draw(screen)
}
