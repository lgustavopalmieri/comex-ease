package usecase

import "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/repository"

type usecase struct {
	repository repository.NcmRepository
}

func NewNcmUseCase(repository repository.NcmRepository) NcmUseCaseInterface {
	return &usecase{
		repository: repository,
	}
}
