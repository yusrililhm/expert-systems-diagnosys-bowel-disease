package repo

import (
	"usus-sehat/server/entity"
	"usus-sehat/server/exception"
)

type DiseaseRepo interface {
	FetchAllDiseases() ([]*entity.Disease, exception.Exception)
}
