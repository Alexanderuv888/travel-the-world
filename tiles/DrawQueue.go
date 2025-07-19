package tiles

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawQueue struct {
	items []Drawable
}

func (dq *DrawQueue) Add(d Drawable) {
	dq.items = append(dq.items, d)
}

func (dq *DrawQueue) Clear() {
	dq.items = dq.items[:0]
}

func (dq *DrawQueue) DrawAll(screen *ebiten.Image, cameraX, cameraY float64) {
	// Сортировка по координате Y
	sort.SliceStable(dq.items, func(i, j int) bool {
		return dq.items[i].ScreenY() < dq.items[j].ScreenY()
	})

	// Отрисовка в порядке глубины
	for _, d := range dq.items {
		d.Draw(screen, cameraX, cameraY)
	}
}
