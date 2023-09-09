package ncmhandler

import (
	"encoding/json"
	"net/http"

	dto "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/dto/create"
)

func (handler ncmhandler) Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	ncmRequest, err := dto.JSONCreateNcmInputDto(request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	ncm, err := handler.usecase.Create(ncmRequest)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(response).Encode(ncm)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
