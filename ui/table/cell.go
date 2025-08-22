package table

import (
	"image"
	"travel-the-world/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Cell struct {
	d        image.Rectangle
	Content  string
	style    Style
	Hovered  bool
	Selected bool
}

func (c *Cell) Update() {
	c.Selected = false

	if c.hovered() {
		c.Hovered = true
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			c.Selected = true
		}
	}
}
func (c *Cell) hovered() bool {
	cx, cy := ebiten.CursorPosition()
	mp := image.Point{X: cx, Y: cy}
	return mp.In(c.d)
}

func (c *Cell) Draw(dst *ebiten.Image) {
	// Draw the cell background
	cellRect := ebiten.NewImage(c.d.Dx(), c.d.Dy())
	cellRect.Fill(c.Style().BackgroundColor)

	// Draw the text inside the cell
	ui.DrawText(cellRect, c.Content, c.Style().TextColor, float64(c.Style().FontSize), 5, c.style.TextAlign)

	// Create draw options and set position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.d.Min.X), float64(c.d.Min.Y))

	// Draw the cell on the destination image
	dst.DrawImage(cellRect, op)
}

func (c *Cell) Style() BaseStyle {
	if c.Hovered {
		return c.style.hs
	}
	if c.Selected {
		return c.style.ss
	}
	return c.style.cs
}
