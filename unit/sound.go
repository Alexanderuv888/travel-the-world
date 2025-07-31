package unit

import (
	"log"
	"math/rand/v2"
	"travel-the-world/assets"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	SwordAttack_0 = "assets/resources/sounds/sword/hit/sword_attack_0.wav"
	SwordAttack_1 = "assets/resources/sounds/sword/hit/sword_attack_1.wav"
	SwordAttack_2 = "assets/resources/sounds/sword/hit/sword_attack_2.wav"
	SwordAttack_3 = "assets/resources/sounds/sword/hit/sword_attack_3.wav"
	death_man_1   = "assets/resources/sounds/male/death/15.wav"
	death_man_2   = "assets/resources/sounds/male/death/16.wav"
	attack_man_1  = "assets/resources/sounds/male/attack/1.wav"
	attack_man_2  = "assets/resources/sounds/male/attack/2.wav"
	attack_man_3  = "assets/resources/sounds/male/attack/3.wav"
	attack_man_4  = "assets/resources/sounds/male/attack/4.wav"
	attack_man_5  = "assets/resources/sounds/male/attack/5.wav"
	death_female  = "assets/resources/sounds/famale/mixkit-female-astonished-gasp-964.wav"
	Melee_1       = "assets/resources/sounds/melee/sword sound.wav"
	Wellmet       = "assets/resources/sounds/male/male-test1_0/wellmet1.wav"
)

type sound struct {
	attackSwordPool   []*audio.Player
	confirmSpeechPool []*audio.Player
	attackSpeechPool  []*audio.Player
	diePool           []*audio.Player
	attackPool        []*audio.Player
}

func NewSound(am *assets.Manager) *sound {
	asp := createPlayerPool(am, SwordAttack_0, SwordAttack_1, SwordAttack_2, SwordAttack_3)
	csp := createPlayerPool(am, Wellmet)
	dp := createPlayerPool(am, death_man_1, death_man_2)
	ap := createPlayerPool(am, attack_man_1, attack_man_2, attack_man_3, attack_man_4, attack_man_5)
	return &sound{
		attackSwordPool:   asp,
		confirmSpeechPool: csp,
		diePool:           dp,
		attackPool:        ap,
	}
}

func createPlayerPool(am *assets.Manager, pathes ...string) []*audio.Player {
	var pp []*audio.Player
	for _, path := range pathes {
		p, err := am.CreatePlyerFor(path)
		if err != nil {
			log.Fatalf("Could not create player for %s error: %v", path, err)
		}
		pp = append(pp, p)
	}
	return pp
}

func (u *Unit) play(path string, volume float64) {
	player, err := u.am.CreatePlyerFor(path)

	if err != nil {
		log.Printf("ошибка загрузки звука, %s: %v", path, err)
		return
	}

	if err := player.Rewind(); err != nil {
		log.Printf("ошибка Rewind: %v", err)
		return
	}

	player.SetVolume(volume)
	player.Play()
}

func (u *Unit) PlayConfirmSpeach() {
	playSound(u.sound.confirmSpeechPool, 1)
}

func (u *Unit) playSwordAttack() {
	playSound(u.sound.attackSwordPool, 0.2)
}

func playSound(pool []*audio.Player, volume float64) {
	i := rand.IntN(len(pool))
	player := pool[i]

	if err := player.Rewind(); err != nil {
		log.Printf("ошибка Rewind: %v", err)
		return
	}
	player.SetVolume(volume)
	player.Play()
}
