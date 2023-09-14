package usecase

import (
	entity "github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/entity"
	"github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/repository"
)

type NcmUsecaseInterface interface {
	// GetById(id string) (entity.NcmInterface, error)
	Create(code, description, startDate, endDate, actType string, actNumber, actYear int) (entity.NcmInterface, error)
}

// NcmService ?
type NcmUseCase struct {
	Persistence repository.NcmPersistenceInterface
}

func NewNcmUseCase(persistence repository.NcmPersistenceInterface) *NcmUseCase {
	return &NcmUseCase{Persistence: persistence}
}
