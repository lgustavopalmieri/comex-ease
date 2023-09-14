package di

import (
	"database/sql"

	"github.com/lgustavopalmieri/comex-ease/adapters/mysqldb/ncmrepositorydb"
	"github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/usecase"
)

func ncmDI(db *sql.DB) *usecase.NcmUseCase {
	ncmDb := ncmrepositorydb.NewNcmRepositoryDb(db)
	ncmUsecase := usecase.NcmUseCase{Persistence: ncmDb}
	return &ncmUsecase
}

func SetupNcmUsecase(db *sql.DB) *usecase.NcmUseCase {
	return ncmDI(db)
}
