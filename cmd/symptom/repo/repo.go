package repo

import (
	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/pkg/exception"
)

type SymptomRepo interface {
	FindAllSymptoms() ([]*entity.Symptom, exception.Exception)
}
