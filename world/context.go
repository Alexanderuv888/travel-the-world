package world

import (
	"travel-the-world/camera"
	"travel-the-world/common"
)

type Context struct {
	InteractableList common.InteractableList
	Camera           *camera.Camera
}

func (ctx *Context) Update(InteractableList common.InteractableList) {
	ctx.InteractableList = InteractableList
}
