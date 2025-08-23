package table

import (
	"image/color"
	text "travel-the-world/ui/text"
)

type BaseStyle struct {
	BackgroundColor color.Color
	TextColor       color.Color
	FontSize        int
}

type Style struct {
	cs BaseStyle // common style
	hs BaseStyle // hovered style
	ss BaseStyle // selected style

	TextAlign text.Alignment // "left", "center", "right"
}

func DefaultStyle() *Style {
	cs := BaseStyle{
		BackgroundColor: color.RGBA{50, 50, 50, 255},
		TextColor:       color.White,
		FontSize:        20}
	hs := BaseStyle{
		BackgroundColor: color.RGBA{60, 60, 60, 255},
		TextColor:       color.White,
		FontSize:        20}
	ss := BaseStyle{
		BackgroundColor: color.RGBA{80, 80, 80, 255},
		TextColor:       color.White,
		FontSize:        20}
	return &Style{
		cs:        cs,
		hs:        hs,
		ss:        ss,
		TextAlign: text.Left,
	}
}

func DefaultStyle2() *Style {
	s := DefaultStyle()
	s.cs.BackgroundColor = color.RGBA{70, 70, 70, 255}
	return s
}

func (s *Style) GetTextSize() int {
	return s.cs.FontSize
}

func (s *Style) SetCommonStyle(style BaseStyle) {
	s.cs = style
}

func (s *Style) SetHoveredStyle(style BaseStyle) {
	s.hs = style
}

func (s *Style) SetSelectedStyle(style BaseStyle) {
	s.ss = style
}
