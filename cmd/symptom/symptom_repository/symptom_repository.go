package symptom_repository

import (
	"healthy-bowel/internal/domain"
	"healthy-bowel/internal/pkg/errors"
)

type SymptomRepository interface {
	Add(symptom *domain.Symptom) errors.Errors
	Delete(id uint) errors.Errors
	Edit(symptom *domain.Symptom) errors.Errors
	GetById(id uint) (*domain.Symptom, errors.Errors)
	GetAll() ([]*domain.Symptom, errors.Errors)
}
