package ncmrepository

import "database/sql"

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	return &MysqlRepository{db: db}
}
