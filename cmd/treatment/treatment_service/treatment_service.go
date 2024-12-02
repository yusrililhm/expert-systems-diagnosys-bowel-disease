package treatment_service

import (
	"healthy-bowel/internal/pkg/dto"
	"healthy-bowel/internal/pkg/errors"
)

type TreatmentService interface {
	Add(payload *dto.TreatmentPayload) errors.Errors
	GetAll() errors.Errors
	GetById(id uint) errors.Errors
	Edit(id uint, payload *dto.TreatmentPayload) errors.Errors
	Delete(id uint) errors.Errors
}
