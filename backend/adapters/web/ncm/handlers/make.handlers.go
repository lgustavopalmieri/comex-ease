package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/usecase"
)

func MakeNcmHandlers(r chi.Router, usecase usecase.NcmUsecaseInterface) {
	r.Route("/ncm", func(r chi.Router) {
		r.Post("/", createNcm(usecase))
	})
}
