package create

import (
	"encoding/json"
	"io"
)

type CreateNcmInputDto struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ActType     string `json:"act_type"`
	ActNumber   int    `json:"act_number"`
	ActYear     int    `json:"act_year"`
}

func JSONCreateNcmInputDto(body io.Reader) (*CreateNcmInputDto, error) {
	request := CreateNcmInputDto{}
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}
