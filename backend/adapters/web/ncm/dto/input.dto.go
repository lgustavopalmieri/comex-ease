package dto

import (
	entity "github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/entity"
)

type NcmInputDto struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ActType     string `json:"act_type"`
	ActNumber   int    `json:"act_number"`
	ActYear     int    `json:"act_year"`
}

func NewNcmInputDto() *NcmInputDto {
	return &NcmInputDto{}
}

func (dto *NcmInputDto) Bind(ncm *entity.Ncm) (*entity.Ncm, error) {
	if dto.ID == "" {
		ncm.ID = dto.ID
	}
	ncm.Code = dto.Code
	ncm.Description = dto.Description
	ncm.StartDate = dto.StartDate
	ncm.EndDate = dto.EndDate
	ncm.ActType = dto.ActType
	ncm.ActNumber = dto.ActNumber
	ncm.ActYear = dto.ActYear
	_, err := ncm.IsValid()
	if err != nil {
		return &entity.Ncm{}, err
	}
	return ncm, nil
}
