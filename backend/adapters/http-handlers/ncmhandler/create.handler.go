package ncmhandler

import (
	"encoding/json"
	"net/http"

	dto "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/dto/create"
	entity "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/entity"
)

func SerializeResponse(responseData entity.Ncm) ([]byte, error) {
	response := entity.Ncm{
		ID:          responseData.ID,
		Code:        responseData.Code,
		Description: responseData.Description,
		StartDate:   responseData.StartDate,
		EndDate:     responseData.EndDate,
		ActType:     responseData.ActType,
		ActNumber:   responseData.ActNumber,
		ActYear:     responseData.ActYear,
		CreatedAt:   responseData.CreatedAt,
		UpdatedAt:   responseData.UpdatedAt,
		DeletedAt:   responseData.DeletedAt,
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (handler ncmhandler) Create(w http.ResponseWriter, r *http.Request) {
	ncmRequest, err := dto.JSONCreateNcmInputDto(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ncm, err := handler.usecase.Create(ncmRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData, err := SerializeResponse(*ncm)
	if err != nil {
		handleHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
	println(responseData)
}

func handleHTTPError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}
