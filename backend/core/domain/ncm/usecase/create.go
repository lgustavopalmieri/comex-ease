package usecase

import (
	"github.com/lgustavopalmieri/comex-ease/core/domain/ncm/dto/create"
	entity "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/entity"

)

func (ncmusecase usecase) Create(ncmRequest *create.CreateNcmInputDto) (*entity.Ncm, error) {
	ncm, err := ncmusecase.repository.Create(ncmRequest)
	if err != nil {
		return nil, err
	}
	return ncm, nil
}
