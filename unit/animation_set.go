package unit

import "fmt"

type USet struct {
	uType  UnitType
	armor  Armor
	head   Head
	weapon Weapon
	shield Shield
}

type UnitType string
type Armor string
type Head string
type Weapon string
type Shield string

func (us *USet) armorTSN() string {
	return fmt.Sprintf("%s/%s", us.uType, us.armor)
}
func (us *USet) headTSN() string {
	return fmt.Sprintf("%s/%s", us.uType, us.head)
}
func (us *USet) weaponTSN() string {
	return fmt.Sprintf("%s/%s", us.uType, us.weapon)
}
func (us *USet) shieldTSN() string {
	return fmt.Sprintf("%s/%s", us.uType, us.shield)
}

func (us *USet) getAllTsn() []string {
	return []string{us.armorTSN(), us.headTSN(), us.weaponTSN(), us.shieldTSN()}
}
