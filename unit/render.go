package unit

import (
	"fmt"
	"image"
	"travel-the-world/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

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
