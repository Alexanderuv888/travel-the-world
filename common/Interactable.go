package common

import "image"

type Interactable interface {
	Target
	ObjType() ObjType
	Rect() image.Rectangle
	Interact(obj Interactable)
}

type ObjType string

const (
	NPC      ObjType = "npc"
	Unit     ObjType = "unit"
	Obstical ObjType = "obstical"
)
