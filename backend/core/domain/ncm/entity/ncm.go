package ncm

import (
	"time"
)

type Ncm struct {
	ID          string     `json:"id"`
	Code        string     `json:"code"`
	Description string     `json:"description"`
	StartDate   string     `json:"start_date"`
	EndDate     string     `json:"end_date"`
	ActType     string     `json:"act_type"`
	ActNumber   int        `json:"act_number"`
	ActYear     int        `json:"act_year"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
