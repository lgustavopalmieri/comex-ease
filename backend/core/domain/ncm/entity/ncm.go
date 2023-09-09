package ncm

import (
	"github.com/lgustavopalmieri/comex-ease/pkg/entity"
)

type Ncm struct {
	ID          entity.ID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	ActType     string    `json:"act_type"`
	ActNumber   int       `json:"act_number"`
	ActYear     int       `json:"act_year"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   *string   `json:"deleted_at,omitempty"`
}

func NewNcm(
	code, description,
	startDate, endDate,
	actType string,
	actNumber, actYear int,
) (*Ncm, error) {
	ncm := &Ncm{
		ID:          entity.NewID(),
		Code:        code,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		ActType:     actType,
		ActNumber:   actNumber,
		ActYear:     actYear,
	}
	return ncm, nil
}
