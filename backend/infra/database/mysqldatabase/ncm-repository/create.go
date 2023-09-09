package ncmrepository

import (
	dto "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/dto/create"
	entity "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/entity"
)

func (repository *MysqlRepository) Create(ncmRequest *dto.CreateNcmInputDto) (*entity.Ncm, error) {
	ncm := &entity.Ncm{}

	query := `INSERT INTO ncms (id, code, description, start_date, end_date, act_type, act_number, act_year)
	VALUES (?,?, ?, ?, ?, ?, ?, ?)`

	stmt, err := repository.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		ncmRequest.ID,
		ncmRequest.Code,
		ncmRequest.Description,
		ncmRequest.StartDate,
		ncmRequest.EndDate,
		ncmRequest.ActType,
		ncmRequest.ActNumber,
		ncmRequest.ActYear,
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
