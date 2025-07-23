package unit

import (
	"travel-the-world/tiles"
)

func (u *Unit) Render(dq *tiles.DrawQueue) {
	u.updateAnimation()
	dq.Add(u)
}
