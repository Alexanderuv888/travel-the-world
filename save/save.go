package save

type SaveData struct {
	Health              int     `json:"Health"`
	MaxHealth           int     `json:"MaxHealth"`
	HealthSpeedRecover  int     `json:"HealthSpeedRecover"`
	Stamina             int     `json:"Stamina"`
	MaxStamina          int     `json:"MaxStamina"`
	StaminaSpeedRecover int     `json:"StaminaSpeedRecover"`
	Damage              int     `json:"damage"`
	AttackDistance      float64 `json:"AttackDistance"`
	Vision              float64 `json:"Vision"`
	Speed               float64 `json:"Speed"`
	Level               int     `json:"Level"`
	MaxExp              int     `json:"MaxExp"`
	Exp                 int     `json:"Exp"`
}
