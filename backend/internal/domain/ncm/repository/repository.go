package repository

import entity "github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/entity"

type NcmReaderInterface interface {
	// GetById(id string) (entity.NcmInterface, error)
}

type NcmWriterInterface interface {
	Create(ncm entity.NcmInterface) (entity.NcmInterface, error)
}

type NcmPersistenceInterface interface {
	NcmReaderInterface
	NcmWriterInterface
}
