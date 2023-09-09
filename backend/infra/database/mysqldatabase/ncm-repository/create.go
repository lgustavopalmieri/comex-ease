package ncmrepository

import (
	"github.com/google/uuid"
	dto "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/dto/create"
	entity "github.com/lgustavopalmieri/comex-ease/core/domain/ncm/entity"
)

func (repository *MysqlRepository) Create(ncmRequest *dto.CreateNcmInputDto) (*entity.Ncm, error) {
	query := `INSERT INTO ncms (id, code, description, start_date, end_date, act_type, act_number, act_year)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := repository.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	ncmID := uuid.New()

	_, err = stmt.Exec(
		ncmID,
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

	row := repository.db.QueryRow(`SELECT * FROM ncms WHERE id = ?`, ncmID)

	var ncm entity.Ncm
	err = row.Scan(
		&ncm.ID,
		&ncm.Code,
		&ncm.Description,
		&ncm.StartDate,
		&ncm.EndDate,
		&ncm.ActType,
		&ncm.ActNumber,
		&ncm.ActYear,
		&ncm.CreatedAt,
		&ncm.UpdatedAt,
		&ncm.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &ncm, nil
}
