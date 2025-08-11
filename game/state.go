package game

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/goregular"
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
)

func (g *Game) Update() error {
	switch g.state {
	case StateMenu:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			x, y := ebiten.CursorPosition()
			if x > 764 && x < 864 && y > 200 && y < 250 { // Играть
				g.state = StatePlaying
			} else if x > 764 && x < 864 && y > 270 && y < 320 { // Выйти
				return ebiten.Termination
			}
		}
	case StatePlaying:
		g.UpdateGame()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateMenu:
		screen.Fill(color.RGBA{30, 30, 30, 255})
		x := screen.Bounds().Dx()/2 - 50

		// Кнопка "Играть"
		vector.DrawFilledRect(screen, float32(x-50), 200, 200, 50, color.RGBA{100, 200, 100, 255}, false)
		drawText(screen, "Играть", x+10, 210, color.White)

		// Кнопка "Выйти"
		vector.DrawFilledRect(screen, float32(x-50), 270, 200, 50, color.RGBA{200, 100, 100, 255}, false)
		drawText(screen, "Выйти", x+15, 280, color.White)

	case StatePlaying:
		g.DrawGame(screen)

	}
}

// Вспомогательная функция для рисования текста
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

	/*text.Draw(screen, str, &text.GoTextFace{
		Source: &text.GoTextFaceSource{
			Face: basicfont.Face7x13, // прямо сюда
		},
		Size: 13, // размер в px
	}, op)*/
}
