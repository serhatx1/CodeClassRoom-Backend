package Models

type TestCase struct {
	ID        uint   `gorm:"primaryKey"`
	ProblemID uint   `json:"problem_id"`
	Input     string `json:"input"`
	Expected  string `json:"expected"`
}
