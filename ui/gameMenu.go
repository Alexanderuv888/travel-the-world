package ui

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type Button struct {
	Text     string
	X, Y     int
	Width    int
	Height   int
	Hovered  bool
	Callback func()
}

type Menu struct {
	Buttons []*Button
}

func NewMenu() *Menu {
	startBtn := &Button{
		Text:   "New Game",
		X:      100,
		Y:      100,
		Width:  200,
		Height: 40,
		Callback: func() {
			if onStart != nil {
				onStart()
			}
		},
	}
	settingsBtn := &Button{
		Text:   "Settings",
		X:      100,
		Y:      160,
		Width:  200,
		Height: 40,
		Callback: func() {
			if onSettings != nil {
				onSettings()
			}
		},
	}
	exitBtn := &Button{
		Text:   "Exit",
		X:      100,
		Y:      220,
		Width:  200,
		Height: 40,
		Callback: func() {
			if onExit != nil {
				onExit()
			}
		},
	}
	return &Menu{Buttons: []*Button{startBtn, settingsBtn, exitBtn}}
}

var (
	onStart    func()
	onSettings func()
	onExit     func()
)

func (m *Menu) SetStartCallback(cb func()) {
	onStart = cb
}

func (m *Menu) SetSettingsCallback(cb func()) {
	onSettings = cb
}

func (m *Menu) SetExitCallback(cb func()) {
	onExit = cb
}

func (m *Menu) Update() {
	x, y := ebiten.CursorPosition()
	for _, b := range m.Buttons {
		b.Hovered = x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height
		if b.Hovered && inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			b.Callback()
		}
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for _, b := range m.Buttons {
		clr := color.RGBA{100, 100, 100, 255}
		if b.Hovered {
			clr = color.RGBA{150, 150, 150, 255}
		}
		sub := ebiten.NewImage(b.Width, b.Height)
		sub.Fill(clr)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(b.X), float64(b.Y))
		screen.DrawImage(sub, op)
		drawText(screen, b.Text, b.X+10, b.Y+6, color.White)
	}
}

func drawText(screen *ebiten.Image, str string, x, y int, clr color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(clr)

	src, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal("Error during creating font. ", err)
	}
	face := text.GoTextFace{
		Source: src,
		Size:   24, // размер шрифта
	}

	text.Draw(
		screen, // *ebiten.Image – изображение кнопки
		str,    // сам текст
		&face,  // интерфейс Face, описывающий шрифт
		op,     // опции рисования, например центровка
	)
}
