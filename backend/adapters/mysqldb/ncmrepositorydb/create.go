package ncmrepositorydb

import (
	entity "github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/entity"
)

func (n *NcmRepositoryDb) Create(ncm entity.NcmInterface) (entity.NcmInterface, error) {
	query := `INSERT INTO ncms (id, code, description, start_date, end_date, act_type, act_number, act_year)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := n.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		ncm.GetID(),
		ncm.GetCode(),
		ncm.GetDescription(),
		ncm.GetStartDate(),
		ncm.GetEndDate(),
		ncm.GetActType(),
		ncm.GetActNumber(),
		ncm.GetActYear(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return ncm, nil
}
