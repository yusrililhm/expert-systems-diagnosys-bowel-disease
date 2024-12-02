package symptom_service

import (
	"healthy-bowel/internal/domain"
	"healthy-bowel/internal/pkg/dto"
	"healthy-bowel/internal/pkg/errors"
)

type SymptomService interface {
	Add(payload *dto.SymptomPayload) errors.Errors
	Delete(id uint) errors.Errors
	Edit(id uint, payload *dto.SymptomPayload) errors.Errors
	GetById(id uint) (*domain.Symptom, errors.Errors)
	GetAll() ([]*domain.Symptom, errors.Errors)
}
