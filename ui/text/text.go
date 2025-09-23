package ui

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type Alignment int

const (
	Left Alignment = iota
	Center
	Right
)

func DrawTextWithCoordinates(dst *ebiten.Image, x, y int, str string, clr color.Color, fontSize float64) {
	src, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal("Error during creating font. ", err)
	}
	face := text.GoTextFace{
		Source: src,
		Size:   fontSize, // размер шрифта
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(clr)

	text.Draw(
		dst,   // *ebiten.Image – бэкграунд, на который рисуем
		str,   // сам текст
		&face, // интерфейс Face, описывающий шрифт
		op,    // опции рисования, например центровка
	)
}

func DrawText(dst *ebiten.Image, str string, clr color.Color, fontSize, offset float64, a Alignment) {
	src, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal("Error during creating font. ", err)
	}
	face := text.GoTextFace{
		Source: src,
		Size:   fontSize, // размер шрифта
	}

	textW, textH := text.Measure(str, &face, float64(5))

	// центрирование внутри dst
	dstW := float64(dst.Bounds().Dx())
	dstH := float64(dst.Bounds().Dy())
	var x float64
	switch a {
	case Left:
		x = offset
	case Center:
		x = (dstW - textW) / 2
	case Right:
		x = dstW - textW - offset
	}
	y := (dstH - textH) / 2

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(clr)

	text.Draw(
		dst,   // *ebiten.Image – бэкграунд, на который рисуем
		str,   // сам текст
		&face, // интерфейс Face, описывающий шрифт
		op,    // опции рисования, например центровка
	)
}
