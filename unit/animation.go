package unit

import (
	"fmt"
	"image"
	"travel-the-world/assets"
	"travel-the-world/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type animation struct {
	aType     assets.Action
	direction assets.Direction
	us        USet
	CTile     *tiles.CompositeTile
	tick      int
	current   int
}

func NewAnimation(x, y float64, am *assets.Manager) *animation {
	us := USet{Hero, Clothes, Male_head1, Shortsword, Buckler}

	armorAnimationSet, _ := am.LoadAnimationSet(unitBaseFolder, us.armorTSN())
	headAnimationSet, _ := am.LoadAnimationSet(unitBaseFolder, us.headTSN())
	weaponAnimationSet, _ := am.LoadAnimationSet(unitBaseFolder, us.weaponTSN())
	shieldAnimationSet, _ := am.LoadAnimationSet(unitBaseFolder, us.shieldTSN())

	armorFrame := armorAnimationSet.GetFrame(fmt.Sprintf("%s/%s", us.armorTSN(), ActionIdle), DirLeft, 0)
	headFrame := headAnimationSet.GetFrame(fmt.Sprintf("%s/%s", us.headTSN(), ActionIdle), DirLeft, 0)
	weaponFrame := weaponAnimationSet.GetFrame(fmt.Sprintf("%s/%s", us.weaponTSN(), ActionIdle), DirLeft, 0)
	shieldFrame := shieldAnimationSet.GetFrame(fmt.Sprintf("%s/%s", us.shieldTSN(), ActionIdle), DirLeft, 0)

	ct := tiles.NewCompositeTile(x, y, frameWidth, frameHeight, armorFrame, headFrame, weaponFrame, shieldFrame)
	return &animation{
		direction: DirLeft,
		aType:     ActionIdle,
		us:        us,
		CTile:     ct,
	}
}

func (u *Unit) updateAnimation(screen *ebiten.Image, levelDimentions image.Point, am *assets.Manager) {
	u.animation.tick++
	armorAS, _ := am.LoadAnimationSet(unitBaseFolder, u.animation.us.armorTSN())
	headAS, _ := am.LoadAnimationSet(unitBaseFolder, u.animation.us.headTSN())
	weaponAS, _ := am.LoadAnimationSet(unitBaseFolder, u.animation.us.weaponTSN())
	shieldAS, _ := am.LoadAnimationSet(unitBaseFolder, u.animation.us.shieldTSN())

	animationName := fmt.Sprintf("%s/%s", u.animation.us.armorTSN(), u.animation.aType)
	if u.animation.tick%armorAS.GetAnimationFrameRate(animationName) == 0 {
		u.animation.current++
		if u.animation.current >= armorAS.GetAnimationLength(animationName, u.animation.direction) {
			u.animation.current = 0
		}
	}
	u.animation.CTile.Tx = u.X - frameWidth*0.5
	u.animation.CTile.Ty = u.Y - frameHeight*0.75

	u.animation.CTile.Images = []*ebiten.Image{armorAS.GetFrame(animationName, u.animation.direction, u.animation.current),
		headAS.GetFrame(fmt.Sprintf("%s/%s", u.animation.us.headTSN(), u.animation.aType), u.animation.direction, u.animation.current),
		weaponAS.GetFrame(fmt.Sprintf("%s/%s", u.animation.us.weaponTSN(), u.animation.aType), u.animation.direction, u.animation.current),
		shieldAS.GetFrame(fmt.Sprintf("%s/%s", u.animation.us.shieldTSN(), u.animation.aType), u.animation.direction, u.animation.current),
	}
}

func (u *Unit) updateAnamationType() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		u.stopUnit()
		u.setAnimation(ActionAttack)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		u.stopUnit()
		u.setAnimation(ActionShoot)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		u.stopUnit()
		u.setAnimation(ActionDie)
	} else if u.vx == 0 || u.vy == 0 {
		u.setAnimation(ActionIdle)
	} else {
		u.setAnimation(ActionRun)
	}
}

func (u *Unit) setAnimation(newAction assets.Action) {
	if u.animation.aType != newAction {
		u.animation.current = 0
		u.animation.aType = newAction
	}
}
