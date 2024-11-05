package repo

import (
	"usus-sehat/server/entity"
	"usus-sehat/server/exception"
)

type TreatmentRepo interface {
	FetchAllTreatments() ([]*entity.Treatment, exception.Exception)
}
