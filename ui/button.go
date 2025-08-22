package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct {
	Text     string
	X, Y     int
	Width    int
	Height   int
	pressed  bool
	Callback func()
}

func NewButton(x, y, width, height int, s string, colbacl func()) *Button {
	return &Button{
		Text:     s,
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
		Callback: colbacl,
	}
}

func (b *Button) Update() {
	if !b.isHovered() {
		b.pressed = false
	}
	if b.isHovered() && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		b.pressed = true
	}
	if b.isHovered() && inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		b.Callback()
		b.pressed = false
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	clr := color.RGBA{100, 100, 100, 255}
	if b.isHovered() {
		clr = color.RGBA{150, 150, 150, 255}
	}
	if b.pressed {
		clr = color.RGBA{50, 50, 50, 255}
	}
	sub := ebiten.NewImage(b.Width, b.Height)
	sub.Fill(clr)

	DrawText(sub, b.Text, color.White, 24, 5, Center)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(sub, op)
}

func (b *Button) isHovered() bool {
	x, y := ebiten.CursorPosition()
	return x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height
}
