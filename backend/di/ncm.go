package di

import (
	"database/sql"

	"github.com/lgustavopalmieri/comex-ease/adapters/http-handlers/ncmhandler"
	usecase "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/usecase"
	ncmrepository "github.com/lgustavopalmieri/comex-ease/infra/database/mysqldatabase/ncm-repository"
)

func ConfigNcmDI(conn *sql.DB) ncmhandler.NcmHandlerInterface {
	ncmRepository := ncmrepository.NewMysqlRepository(conn)
	ncmUsecase := usecase.NewNcmUseCase(ncmRepository)
	ncmHandler := ncmhandler.NewNcmHandler(ncmUsecase)
	return ncmHandler
}
