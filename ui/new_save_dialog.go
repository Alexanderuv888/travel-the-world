package ui

import (
	"image/color"
	text "travel-the-world/ui/text"

	"github.com/hajimehoshi/ebiten/v2"
)

type NewSaveDialog struct {
	X, Y, Width, Height int
	Visible             bool
	newSaveName         string
	buttonsBar          *ButtonsBar
}

func NewNewSaveDialog(X, Y, Width, Height int) *NewSaveDialog {
	d := &NewSaveDialog{
		X:      X,
		Y:      Y,
		Width:  Width,
		Height: Height,
	}
	d.buttonsBar = NewButtonsBar(X+5, Y+Height-50, Width-10, 50)
	d.buttonsBar.SetOrientation(OrientHorizontal)
	d.buttonsBar.SetAlignment(AlignRight)
	d.buttonsBar.SetPadding(10)
	d.buttonsBar.SetSpacing(10)
	return d
}

func (d *NewSaveDialog) Update() {
	if !d.Visible {
		return
	}
	d.buttonsBar.Update()
	// Обновляем строку ввода имени файла
	for _, r := range ebiten.AppendInputChars(nil) {
		if r == '\n' || r == '\r' {
			continue
		}
		d.newSaveName += string(r)
	}
}
func (d *NewSaveDialog) Draw(screen *ebiten.Image) {
	if !d.Visible {
		return
	}
	dlgBackground := ebiten.NewImage(d.Width, d.Height)
	dlgBackground.Fill(color.RGBA{40, 40, 40, 255})

	dlgOp := &ebiten.DrawImageOptions{}
	dlgOp.GeoM.Translate(float64(d.X), float64(d.Y))

	clr := color.RGBA{50, 50, 50, 255} // обычная строка

	text.DrawTextWithCoordinates(dlgBackground, 10,10, "Enter save name:", color.White, 10)
	

	row := ebiten.NewImage(290, 32)
	row.Fill(clr)
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(10), float64(50))
	text.DrawText(row, d.newSaveName, color.RGBA{200, 200, 200, 255}, 20, 5, text.Right)
	screen.DrawImage(row, op2)
	d.buttonsBar.Draw(screen)

	screen.DrawImage(dlgBackground, dlgOp)
}
