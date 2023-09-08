package repository

import (
	"github.com/lgustavopalmieri/comex-ease/core/domain/ncm/dto/create"
	ncm "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/entity"
)

type NcmRepository interface {
	Create(ncmRequest *create.CreateNcmInputDto) (*ncm.Ncm, error)
}
