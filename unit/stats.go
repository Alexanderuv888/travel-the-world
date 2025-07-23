package unit

type stats struct {
	status         Status
	health         int
	maxHealth      int
	damage         int
	attackDistance float64
	vision         float64
	speed          float64
}

type Status string

const (
	Alive Status = "alive"
	Dead  Status = "dead"
)
