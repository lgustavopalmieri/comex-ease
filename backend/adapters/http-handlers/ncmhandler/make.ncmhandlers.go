package ncmhandler

import (
	"net/http"

	"github.com/lgustavopalmieri/comex-ease/core/domain/ncm/usecase"
)

type NcmHandlerInterface interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type ncmhandler struct {
	usecase usecase.NcmUseCaseInterface
}

func NewNcmHandler(usecase usecase.NcmUseCaseInterface) NcmHandlerInterface {
	return &ncmhandler{
		usecase: usecase,
	}
}
