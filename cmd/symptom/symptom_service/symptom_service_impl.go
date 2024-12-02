package symptom_service

import (
	"healthy-bowel/cmd/symptom/symptom_repository"
	"healthy-bowel/internal/domain"
	"healthy-bowel/internal/pkg/dto"
	"healthy-bowel/internal/pkg/errors"

	"sync"
)

type symptomServiceImpl struct {
	sr symptom_repository.SymptomRepository
	wg *sync.WaitGroup
}

// Add implements SymptomService.
func (ss *symptomServiceImpl) Add(payload *dto.SymptomPayload) errors.Errors {
	panic("unimplemented")
}

// Delete implements SymptomService.
func (ss *symptomServiceImpl) Delete(id uint) errors.Errors {
	panic("unimplemented")
}

// Edit implements SymptomService.
func (ss *symptomServiceImpl) Edit(id uint, payload *dto.SymptomPayload) errors.Errors {
	panic("unimplemented")
}

// GetAll implements SymptomService.
func (ss *symptomServiceImpl) GetAll() ([]*domain.Symptom, errors.Errors) {
	panic("unimplemented")
}

// GetById implements SymptomService.
func (ss *symptomServiceImpl) GetById(id uint) (*domain.Symptom, errors.Errors) {
	panic("unimplemented")
}

func NewSymptomServiceImpl(sr symptom_repository.SymptomRepository, wg *sync.WaitGroup) SymptomService {
	return &symptomServiceImpl{
		sr: sr,
		wg: wg,
	}
}
