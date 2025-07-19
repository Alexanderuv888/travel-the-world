package unit

import "travel-the-world/assets"

const (
	frameWidth     = 128
	frameHeight    = 128
	totalFrames    = 4
	frameDelay     = 10 // сколько обновлений на кадр
	speed          = 1.5
	unitBaseFolder = "assets/resources/unit_tiles"
)

const (
	ActionIdle   assets.Action = "idle"
	ActionRun    assets.Action = "run"
	ActionAttack assets.Action = "attack"
	ActionDie    assets.Action = "die"
	ActionShoot  assets.Action = "shoot"

	DirLeft      assets.Direction = "left"
	DirLeftUp    assets.Direction = "leftUp"
	DirUp        assets.Direction = "up"
	DirRightUp   assets.Direction = "rightUp"
	DirRight     assets.Direction = "right"
	DirRightDown assets.Direction = "rightDown"
	DirDown      assets.Direction = "down"
	DirLeftDown  assets.Direction = "leftDown"
)

const (
	Hero    UnitType = "hero"
	Heroine UnitType = "heroine"

	Clothes       Armor = "clothes"
	Leather_armor Armor = "leather_armor"
	Steel_armor   Armor = "steel_armor"

	Male_head1 Head = "male_head1"
	Male_head2 Head = "male_head2"
	Male_head3 Head = "male_head3"
	Head_long  Head = "head_long"

	Dagger     Weapon = "dagger"
	Greatbow   Weapon = "greatbow"
	Greatsword Weapon = "greatsword"
	Longbow    Weapon = "longbow"
	Longsword  Weapon = "longsword"
	Rod        Weapon = "rod"
	Shortbow   Weapon = "shortbow"
	Shortsword Weapon = "shortsword"
	Slingshot  Weapon = "slingshot"
	Staff      Weapon = "staff"
	Wand       Weapon = "wand"

	Buckler Shield = "buckler"
	shield  Shield = "shield"
)
