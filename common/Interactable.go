package common

import "image"

type Interactable interface {
	ObjType() ObjType
	Rect() image.Rectangle
	Centr() image.Point
	Interact(obj Interactable)
}

type ObjType string

const (
	NPC      ObjType = "npc"
	Unit     ObjType = "unit"
	Obstical ObjType = "obstical"
)
