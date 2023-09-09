package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lgustavopalmieri/comex-ease/di"
	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/comex-test")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ncmHandlerProvider := di.ConfigNcmDI(conn)

	router := mux.NewRouter()
	router.Handle("/ncm", http.HandlerFunc(ncmHandlerProvider.Create)).Methods("POST")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
