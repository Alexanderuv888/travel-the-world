package game

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePaused
	StateGameOver
	StateQuitGame
)

func (g *Game) Update() error {
	g.updateState()
	switch g.state {
	case StateMenu:
		g.Menu.Update()
	case StatePlaying:
		g.UpdateGame()
	case StateQuitGame:
		return ebiten.Termination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateMenu:
		g.Menu.Draw(screen)
	case StatePlaying:
		g.DrawGame(screen)
	}
}

// Вспомогательная функция для рисования текста
func DrawText(screen *ebiten.Image, str string, x, y int, clr color.Color) {
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

func (g *Game) updateState() {
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		if g.state == StatePlaying {
			g.state = StateMenu
		} else {
			g.state = StatePlaying
		}
	}
}
