package ncm

import (
	"errors"

	"github.com/asaskevich/govalidator"

	id "github.com/lgustavopalmieri/comex-ease/pkg/entity"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Ncm struct {
	ID          string  `valid:"uuidv4"`
	Code        string  `valid:"required"`
	Description string  `valid:"required"`
	StartDate   string  `valid:"required"`
	EndDate     string  `valid:"required"`
	ActType     string  `valid:"required"`
	ActNumber   int     `valid:"required"`
	ActYear     int     `valid:"required"`
	CreatedAt   string  `valid:"required"`
	UpdatedAt   string  `valid:"required"`
	DeletedAt   *string `valid:"string,optional"`
}

func NewNcm(code, description, startDate, endDate, actType string, actNumber, actYear int) *Ncm {
	ncm := Ncm{
		ID:          id.NewID().String(),
		Code:        code,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		ActType:     actType,
		ActNumber:   actNumber,
		ActYear:     actYear,
	}

	return &ncm
}

func (n *Ncm) IsValid() (bool, error) {
	if n.Code == "" {
		return false, errors.New("missing ncm code")
	}
	return true, nil
}
