package table

import (
	"image"
	text "travel-the-world/ui/text"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cell struct {
	X, Y, Width, Height int
	D                   image.Rectangle
	Content             string
	style               *Style
	Hovered             bool
	Selected            bool
}

func NewCell(w, h int, c string, s *Style) *Cell {
	return &Cell{
		Width:   w,
		Height:  h,
		Content: c,
		style:   s,
	}
}

func (c *Cell) Update() {

}

func (c *Cell) Draw(dst *ebiten.Image) {
	// Draw the cell background
	cellRect := ebiten.NewImage(c.Width, c.Height)
	cellRect.Fill(c.Style().BackgroundColor)

	// Draw the text inside the cell
	text.DrawText(cellRect, c.Content, c.Style().TextColor, float64(c.Style().FontSize), 5, c.style.TextAlign)

	// Create draw options and set position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y))

	// Draw the cell on the destination image
	dst.DrawImage(cellRect, op)
}

func (c *Cell) Style() BaseStyle {
	if c.Selected {
		return c.style.ss
	}
	if c.Hovered {
		return c.style.hs
	}
	return c.style.cs
}
