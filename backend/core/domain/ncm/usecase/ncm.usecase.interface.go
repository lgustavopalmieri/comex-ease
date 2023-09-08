package usecase

import (
	ncm "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/entity"
	"github.com/lgustavopalmieri/comex-ease/core/domain/ncm/dto/create"
)

type NcmUseCaseInterface interface {
	Create(ncmRequest *create.CreateNcmInputDto) (*ncm.Ncm, error)
}
