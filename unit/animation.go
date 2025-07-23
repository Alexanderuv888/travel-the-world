package unit

import (
	"fmt"
	"image/color"
	"log"
	"travel-the-world/assets"
	"travel-the-world/tiles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type animation struct {
	aSets     map[string]*assets.AnimationSet
	aType     assets.Action
	direction assets.Direction
	us        USet
	CTile     *tiles.CompositeTile
	tick      int
	current   int
	len       int
	rate      int
	freez     bool
}

func NewAnimation(x, y float64, am *assets.Manager) *animation {
	us := USet{Hero, Clothes, Male_head1, Shortsword, Buckler}

	a := &animation{
		aSets:     make(map[string]*assets.AnimationSet),
		direction: DirLeft,
		aType:     ActionIdle,
		us:        us,
	}
	a.loadAS(am)
	ct := tiles.NewCompositeTile(x, y, frameWidth, frameHeight, a.getImages())
	a.CTile = ct
	a.updateASLen()
	a.updateASRate()
	return a
}

func (u *Unit) Width() float64 {
	return u.action.animation.CTile.Width()
}

func (u *Unit) Height() float64 {
	return u.action.animation.CTile.Height()
}

func (u *Unit) ScreenY() float64 {
	return u.action.animation.CTile.ScreenY()
}

func (u *Unit) Draw(screen *ebiten.Image, cameraX, cameraY float64) {
	if u.IsAlive() {
		u.drawHealth(screen, cameraX, cameraY)
	}
	u.action.animation.CTile.Draw(screen, cameraX, cameraY)
}

func (u *Unit) drawHealth(screen *ebiten.Image, cameraX, cameraY float64) {
	lineLenght := 20
	lineHeight := 60
	mx1 := u.Point().X - lineLenght/2 - int(cameraX)
	mx2 := mx1 + lineLenght
	my1 := u.Point().Y - lineHeight - int(cameraY)
	my2 := my1
	vector.StrokeLine(screen, float32(mx1), float32(my1), float32(mx2), float32(my2), 2, color.RGBA{255, 0, 0, 255}, false)

	x1 := u.Point().X - lineLenght/2 - int(cameraX)
	x2 := x1 + ((lineLenght / u.stats.maxHealth) * u.stats.health)
	y1 := u.Point().Y - lineHeight - int(cameraY)
	y2 := y1
	vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 2, color.RGBA{0, 255, 0, 255}, false)
}

func (u *Unit) updateAnimation() {
	if u.action.animation.freez {
		return
	}
	u.action.animation.tick++

	if u.action.animation.tick%u.action.animation.rate == 0 {
		u.action.animation.current++
		if u.action.animation.current >= u.action.animation.len {
			u.action.animation.current = 0
		}
	}
	u.action.animation.CTile.Tx = u.X - frameWidth*0.5
	u.action.animation.CTile.Ty = u.Y - frameHeight*0.75

	u.action.animation.CTile.Images = u.action.animation.getImages()
}

func (u *Unit) setAnimation(newAction assets.Action, needToStop bool) {
	if u.action.animation.aType != newAction {
		if needToStop {
			u.stopUnit()
		}
		u.action.animation.current = 0
		u.action.animation.aType = newAction
		u.action.animation.loadAS(u.am)
		u.action.animation.updateASLen()
		u.action.animation.updateASRate()
	}
}

func (a *animation) getImages() (images []*ebiten.Image) {
	for name, as := range a.aSets {
		animName := fmt.Sprintf("%s/%s", name, a.aType)
		image := as.GetFrame(animName, a.direction, a.current)
		if image != nil {
			images = append(images, image)
		}
	}
	return images
}

func (a *animation) updateASLen() {
	a.len = a.aSets[a.us.armorTSN()].GetAnimationLength(fmt.Sprintf("%s/%s", a.us.armorTSN(), a.aType), a.direction)
}

func (a *animation) updateASRate() {
	animationName := fmt.Sprintf("%s/%s", a.us.armorTSN(), a.aType)
	a.rate = a.aSets[a.us.armorTSN()].GetAnimationFrameRate(animationName)

}

func (a *animation) loadAS(am *assets.Manager) {
	for _, tsn := range a.us.getAllTsn() {
		as, err := am.LoadAnimationSet(unitBaseFolder, tsn)
		if err != nil {
			log.Fatal(err)
		} else {
			a.aSets[string(tsn)] = as
		}
	}
}
