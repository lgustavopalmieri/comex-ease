package ncmrepositorydb

import "database/sql"

type NcmRepositoryDb struct {
	db *sql.DB
}

func NewNcmRepositoryDb(db *sql.DB) *NcmRepositoryDb {
	return &NcmRepositoryDb{db: db}
}
