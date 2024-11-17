package repo

import (
	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/pkg/exception"
)

type TreatmentRepo interface {
	FetchAllTreatments() ([]*entity.Treatment, exception.Exception)
}
