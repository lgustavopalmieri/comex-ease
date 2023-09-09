package mysqldatabase

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func GetMysqlConnection(context context.Context) *sql.DB {
	credentials := viper.GetString("database.url")
	conn, err := sql.Open("mysql", credentials)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
