package unit

import (
	"image"
	"travel-the-world/common"
)

func (u *Unit) faceWithObjects(objects *common.InteractableList) bool {
	for _, o := range objects.Items {
		if u != o && u.Rect().Overlaps(o.Rect()) {
			overlapSize := u.Centr().Sub(o.Centr())
			u.X = float64(u.Centr().X + overlapSize.X/4)
			u.Y = float64(u.Centr().Y + overlapSize.Y/4)
			if u.Goal().In(o.Rect()) {
				u.stopUnit()
			}
			return true
		}
	}
	return false
}

func (u *Unit) ObjType() common.ObjType {
	return common.Unit
}

func (u *Unit) Interact(obj common.Interactable) {
	if u.Rect().Overlaps(obj.Rect()) {
		obj.Rect().Sub(u.Centr())
	}
}

func (u *Unit) Rect() image.Rectangle {
	return image.Rect(int(u.X-(frameWidth/8)), int(u.Y-(frameHeight/8)), int(u.X+(frameWidth/8)), int(u.Y+(frameHeight/8)))
}

func centrOfRect(r image.Rectangle) image.Point {
	return image.Point{(r.Max.X + r.Min.X) / 2, (r.Max.Y + r.Min.Y) / 2}
}

func (u *Unit) Centr() image.Point {
	return image.Point{int(u.X), int(u.Y)}
}

func (u *Unit) Goal() image.Point {
	return image.Point{int(u.GoalX), int(u.GoalY)}
}
