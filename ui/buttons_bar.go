package ui

import "github.com/hajimehoshi/ebiten/v2"

type Orientation int
type Alignment int

const (
	OrientHorizontal Orientation = iota
	OrientVertical
)

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
	AlignTop  = AlignLeft  // alias for readability when used with vertical orientation
	AlignDown = AlignRight // alias for readability when used with vertical orientation
)

type ButtonsBar struct {
	Buttons []*Button
	X, Y    int
	Width   int
	Height  int
	visible bool

	Orientation Orientation
	Alignment   Alignment

	Padding int // inner padding from edges
	Spacing int // explicit spacing; if 0 — computed automatically
}

func NewButtonsBar(X, Y, Width, Height int) *ButtonsBar {
	initButtons()
	return &ButtonsBar{
		X:           X,
		Y:           Y,
		Width:       Width,
		Height:      Height,
		Orientation: OrientHorizontal,
		Alignment:   AlignRight,
		Padding:     8,
		Spacing:     0,
	}
}

func (bb *ButtonsBar) SetVisible(visible bool) {
	bb.visible = visible
}

func (bb *ButtonsBar) SetOrientation(o Orientation) {
	bb.Orientation = o
}

func (bb *ButtonsBar) SetAlignment(a Alignment) {
	bb.Alignment = a
}

func (bb *ButtonsBar) SetPadding(p int) {
	bb.Padding = p
}

func (bb *ButtonsBar) SetSpacing(s int) {
	bb.Spacing = s
}

func totalButtonsWidth(buttons []*Button) int {
	total := 0
	for _, btn := range buttons {
		total += btn.Width
	}
	return total
}

func totalButtonsHeight(buttons []*Button) int {
	total := 0
	for _, btn := range buttons {
		total += btn.Height
	}
	return total
}

func (bb *ButtonsBar) SetButtons(buttons []*Button) {
	// Default cross-axis centering
	for _, btn := range buttons {
		if bb.Orientation == OrientHorizontal {
			btn.Y = (bb.Height-btn.Height)/2 + bb.Y
		} else {
			btn.X = (bb.Width-btn.Width)/2 + bb.X
		}
	}

	// Layout along main axis
	if bb.Orientation == OrientHorizontal {
		bb.layoutHorizontal(buttons)
	} else {
		bb.layoutVertical(buttons)
	}

	bb.Buttons = buttons
}

func (bb *ButtonsBar) layoutHorizontal(buttons []*Button) {
	if len(buttons) == 0 {
		return
	}
	totalW := totalButtonsWidth(buttons)
	avail := bb.Width - 2*bb.Padding

	// compute spacing
	spacing := bb.Spacing
	if spacing <= 0 {
		// distribute remaining space between buttons
		if avail > totalW {
			spacing = (avail - totalW) / (len(buttons) + 1)
		} else {
			spacing = 4 // minimal fallback
		}
	}

	var startX int
	switch bb.Alignment {
	case AlignLeft:
		startX = bb.X + bb.Padding
	case AlignCenter:
		startX = bb.X + (bb.Width-totalW)/2
	case AlignRight:
		startX = bb.X + bb.Width - bb.Padding - totalW
	default:
		startX = bb.X + bb.Padding
	}

	// when using spacing-based distribution for left/center alignment, adjust to include spacing before first
	if bb.Alignment == AlignCenter {
		// center ignores spacing calc and centers tightly; to keep gap between buttons use spacing computed above
		// we will place buttons contiguously starting at startX
	} else if bb.Alignment == AlignLeft || bb.Alignment == AlignRight {
		// keep buttons contiguous (startX already set)
	}

	// If using balanced spacing (when alignment is center, emulate spacing between buttons)
	if bb.Alignment == AlignCenter && avail > totalW {
		startX = bb.X + bb.Padding + spacing
	}

	curX := startX
	for i, btn := range buttons {
		// ensure if center alignment we keep approximate spacing between buttons
		if i > 0 {
			curX += spacing
		}
		btn.X = curX
		curX += btn.Width
	}
}

func (bb *ButtonsBar) layoutVertical(buttons []*Button) {
	if len(buttons) == 0 {
		return
	}
	totalH := totalButtonsHeight(buttons)
	avail := bb.Height - 2*bb.Padding

	spacing := bb.Spacing
	if spacing <= 0 {
		if avail > totalH {
			spacing = (avail - totalH) / (len(buttons) + 1)
		} else {
			spacing = 4
		}
	}

	var startY int
	switch bb.Alignment {
	case AlignTop:
		startY = bb.Y + bb.Padding
	case AlignCenter:
		startY = bb.Y + (bb.Height-totalH)/2
	case AlignDown:
		startY = bb.Y + bb.Height - bb.Padding - totalH
	default:
		startY = bb.Y + bb.Padding
	}

	// similar spacing behavior as horizontal
	if bb.Alignment == AlignCenter && avail > totalH {
		startY = bb.Y + bb.Padding + spacing
	}

	curY := startY
	for i, btn := range buttons {
		if i > 0 {
			curY += spacing
		}
		btn.Y = curY
		curY += btn.Height
	}
}

func (bb *ButtonsBar) Update() {
	if !bb.visible {
		return
	}
	cx, cy := ebiten.CursorPosition()
	for _, btn := range bb.Buttons {
		btn.isHovered(cx, cy)
		btn.Update()
	}
}

func (bb *ButtonsBar) Draw(dst *ebiten.Image) {
	if !bb.visible {
		return
	}
	for _, btn := range bb.Buttons {
		btn.Draw(dst)
	}
}

func (bb *ButtonsBar) SetButtons1(buttons []*Button) {
	for _, btn := range buttons {
		btn.Y = (bb.Height-btn.Height)/2 + bb.Y
	}
	// Распределяем кнопки по ширине
	if len(buttons) > 0 {
		spacing := (bb.Width - totalButtonsWidth(buttons)) / (len(buttons) + 1)
		currentX := spacing + bb.X
		for _, btn := range buttons {
			btn.X = currentX
			currentX += btn.Width + spacing
		}
	}
	bb.Buttons = buttons
}
