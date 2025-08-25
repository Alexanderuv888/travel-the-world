package ui

import (
	"image/color"
	text "travel-the-world/ui/text"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct {
	Text     string
	X, Y     int
	Width    int
	Height   int
	pressed  bool
	hovered  bool
	Callback func()
	invoke   *func()
}

func DefButton(s string, invoke *func()) *Button {
	return NewButton(0, 0, 150, 40, s, invoke)
}

func NewButton(X, Y, Width, Height int, s string, invoke *func()) *Button {

	b := &Button{
		Text:   s,
		X:      X,
		Y:      Y,
		Width:  150,
		Height: 40,
		invoke: invoke,
	}
	b.Callback = func() {
		if b.invoke != nil {
			(*b.invoke)()
		}
	}
	return b
}

func (b *Button) Update() {
	if !b.hovered {
		b.pressed = false
	}
	if b.hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		b.pressed = true
	}
	if b.hovered && inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		b.Callback()
		b.pressed = false
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	clr := color.RGBA{100, 100, 100, 255}
	if b.hovered {
		clr = color.RGBA{150, 150, 150, 255}
	}
	if b.pressed {
		clr = color.RGBA{50, 50, 50, 255}
	}
	sub := ebiten.NewImage(b.Width, b.Height)
	sub.Fill(clr)

	text.DrawText(sub, b.Text, color.White, 24, 5, text.Center)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(sub, op)
}

func (b *Button) isHovered(x, y int) bool {
	b.hovered = x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height
	return b.hovered
}
