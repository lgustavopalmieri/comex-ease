package usecase

import entity "github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/entity"

func (n *NcmUseCase) Create(code, description, startDate, endDate, actType string, actNumber, actYear int) (entity.NcmInterface, error) {
	ncm := entity.NewNcm(
		code,
		description,
		startDate,
		endDate,
		actType,
		actNumber,
		actYear,
	)

	_, err := ncm.IsValid()
	if err != nil {
		return &entity.Ncm{}, err
	}
	result, err := n.Persistence.Create(ncm)
	if err != nil {
		return &entity.Ncm{}, err
	}
	return result, nil
}
