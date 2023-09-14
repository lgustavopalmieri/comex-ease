package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lgustavopalmieri/comex-ease/adapters/web/ncm/dto"
	"github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/usecase"
)

func createNcm(usecase usecase.NcmUsecaseInterface) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var ncmDto dto.NcmInputDto
		err := json.NewDecoder(r.Body).Decode(&ncmDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		ncm, err := usecase.Create(
			ncmDto.Code, 
			ncmDto.Description,
			ncmDto.StartDate,
			ncmDto.EndDate,
			ncmDto.ActType,
			ncmDto.ActNumber,
			ncmDto.ActYear,
		)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(ncm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
