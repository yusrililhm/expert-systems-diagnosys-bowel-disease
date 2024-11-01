package repo

import (
	"usus-sehat/server/entity"
	"usus-sehat/server/exception"
)

type SymptomRepo interface {
	FindAllSymptoms() ([]*entity.Symptom, exception.Exception)
}
