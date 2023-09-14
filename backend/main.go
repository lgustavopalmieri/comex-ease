package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lgustavopalmieri/comex-ease/configs"
	"github.com/lgustavopalmieri/comex-ease/di"
	"github.com/lgustavopalmieri/comex-ease/infra/httpserver"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("Erro ao carregar configurações:", err)
		os.Exit(1)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		os.Exit(1)
	}

	server := httpserver.MakeNewWebServer()
	server.NcmUsecase = di.SetupNcmUsecase(db)
	fmt.Println("Webserver has been started")
	server.Serve()
}
