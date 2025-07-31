package world

import (
	"image"
	"travel-the-world/common"
)

type Context struct {
	InteractableList common.InteractableList
	CameraPos        image.Point
}

func (ctx *Context) Update(cameraPos image.Point, InteractableList common.InteractableList) {
	ctx.CameraPos = cameraPos
	ctx.InteractableList = InteractableList
}
