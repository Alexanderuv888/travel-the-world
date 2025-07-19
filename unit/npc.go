package unit

import (
	"fmt"
	"image"
	"travel-the-world/assets"
	"travel-the-world/common"
	"travel-the-world/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type Unit struct {
	us           USet
	Action       assets.Action
	direction    assets.Direction
	current      int
	tick         int
	X, Y         float64 // позиция юнита
	vx, vy       float64 // скорость юнита
	GoalX, GoalY float64 // желаемая позиция юнита
	Angle        float64 // угол траектории движения юнита относительно оси y в радианах
	sx, sy       float64 // масштабирование по x и y
	CTile        *tiles.CompositeTile
	TopAngle     float64
	BottomAngle  float64
	Health       int
}

func NewUnit(x, y float64, am *assets.Manager) *Unit {
	us := USet{Hero, Clothes, Male_head1, Shortsword, Buckler}

	armorAnimationSet, _ := am.LoadAnimationSet(unitBaseFolder, us.armorTSN())
	headAnimationSet, _ := am.LoadAnimationSet(unitBaseFolder, us.headTSN())
	weaponAnimationSet, _ := am.LoadAnimationSet(unitBaseFolder, us.weaponTSN())
	shieldAnimationSet, _ := am.LoadAnimationSet(unitBaseFolder, us.shieldTSN())

	armorFrame := armorAnimationSet.GetFrame(fmt.Sprintf("%s/%s", us.armorTSN(), ActionIdle), DirLeft, 0)
	headFrame := headAnimationSet.GetFrame(fmt.Sprintf("%s/%s", us.headTSN(), ActionIdle), DirLeft, 0)
	weaponFrame := weaponAnimationSet.GetFrame(fmt.Sprintf("%s/%s", us.weaponTSN(), ActionIdle), DirLeft, 0)
	shieldFrame := shieldAnimationSet.GetFrame(fmt.Sprintf("%s/%s", us.shieldTSN(), ActionIdle), DirLeft, 0)

	ut := tiles.NewCompositeTile(x, y, frameWidth, frameHeight, armorFrame, headFrame, weaponFrame, shieldFrame)

	u := &Unit{
		us:        us,
		X:         x,
		Y:         y,
		vx:        0,
		vy:        0,
		Angle:     0,
		direction: DirLeft,
		Action:    ActionIdle,
		sx:        1,
		sy:        1,
		CTile:     ut,
		Health:    100,
	}
	u.GoalX = float64(u.X)
	u.GoalY = float64(u.Y)
	return u
}

func IsoToWorld(x, y float64, levelDimentions image.Point) (isoX, isoY float64) {
	isoX = (x + y) - float64(levelDimentions.X/2)
	isoY = (x - y)
	return
}

func WorldToIso(x, y float64, levelDimentions image.Point) (isoX, isoY float64) {
	isoX = (x - y) + float64(levelDimentions.X/2)
	isoY = (x + y)
	return
}

func (u *Unit) ObjType() common.ObjType {
	return common.Unit
}

func (u *Unit) Interact(obj common.Interactable) {
	if u.Rect().Overlaps(obj.Rect()) {
		obj.Rect().Sub(u.Centr())

	}
}

func (u *Unit) updateAction() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		u.stopUnit()
		u.SetAction(ActionAttack)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		u.stopUnit()
		u.SetAction(ActionShoot)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		u.stopUnit()
		u.SetAction(ActionDie)
	} else if u.vx == 0 || u.vy == 0 {
		u.SetAction(ActionIdle)
	} else {
		u.SetAction(ActionRun)
	}
}

func (u *Unit) stopUnit() {
	u.GoalX = float64(u.X)
	u.GoalY = float64(u.Y)
	u.vx = 0
	u.vy = 0
}

func (u *Unit) SetAction(newAction assets.Action) {
	if u.Action != newAction {
		u.current = 0
		u.Action = newAction
	}
}

func (u *Unit) tryMove(objects *common.InteractableList) {
	u.X += u.vx
	u.Y += u.vy
	if objects != nil && u.faceWithObjects(objects) {
		u.X -= u.vx
	}
	if objects != nil && u.faceWithObjects(objects) {
		u.Y -= u.vy
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

func (u *Unit) Update(screen *ebiten.Image, levelDimentions image.Point, am *assets.Manager) {
	u.tick++
	armorAS, _ := am.LoadAnimationSet(unitBaseFolder, u.us.armorTSN())
	headAS, _ := am.LoadAnimationSet(unitBaseFolder, u.us.headTSN())
	weaponAS, _ := am.LoadAnimationSet(unitBaseFolder, u.us.weaponTSN())
	shieldAS, _ := am.LoadAnimationSet(unitBaseFolder, u.us.shieldTSN())

	animationName := fmt.Sprintf("%s/%s", u.us.armorTSN(), u.Action)
	if u.tick%armorAS.GetAnimationFrameRate(animationName) == 0 {
		u.current++
		if u.current >= armorAS.GetAnimationLength(animationName, u.direction) {
			u.current = 0
		}
	}
	u.CTile.Tx = u.X - frameWidth*0.5
	u.CTile.Ty = u.Y - frameHeight*0.75

	/*x := u.X - frameWidth*0.5
	y := u.Y - frameHeight*0.75

	u.CTile.Tx, u.CTile.Ty = WorldToIso(x, y, levelDimentions)*/

	u.CTile.Images = []*ebiten.Image{armorAS.GetFrame(animationName, u.direction, u.current),
		headAS.GetFrame(fmt.Sprintf("%s/%s", u.us.headTSN(), u.Action), u.direction, u.current),
		weaponAS.GetFrame(fmt.Sprintf("%s/%s", u.us.weaponTSN(), u.Action), u.direction, u.current),
		shieldAS.GetFrame(fmt.Sprintf("%s/%s", u.us.shieldTSN(), u.Action), u.direction, u.current),
	}
}
