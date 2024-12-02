package treatment_repository

import (
	"healthy-bowel/internal/domain"
	"healthy-bowel/internal/pkg/errors"
)

type TreatmentRepository interface {
	Add(treatment *domain.Treatment) errors.Errors
	GetAll() ([]*domain.Treatment, errors.Errors)
	GetById(id uint) (*domain.Treatment, errors.Errors)
	Edit(treatment *domain.Treatment) errors.Errors
	Delete(id uint) errors.Errors
}
