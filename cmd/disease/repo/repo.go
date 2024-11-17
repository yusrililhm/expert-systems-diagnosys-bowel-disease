package repo

import (
	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/pkg/exception"
)

type DiseaseRepo interface {
	FetchAllDiseases() ([]*entity.Disease, exception.Exception)
}
