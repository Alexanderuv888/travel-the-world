package table

import (
	"image/color"
	"travel-the-world/ui"
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

	TextAlign ui.Alignment // "left", "center", "right"
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
