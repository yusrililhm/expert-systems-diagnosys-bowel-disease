package treatment_service

import (
	"healthy-bowel/cmd/treatment/treatment_repository"
	"healthy-bowel/internal/pkg/dto"
	"healthy-bowel/internal/pkg/errors"
	"sync"
)

type treatmentServiceImpl struct {
	tr treatment_repository.TreatmentRepository
	wg *sync.WaitGroup
}

// Add implements TreatmentService.
func (t *treatmentServiceImpl) Add(payload *dto.TreatmentPayload) errors.Errors {
	panic("unimplemented")
}

// Delete implements TreatmentService.
func (t *treatmentServiceImpl) Delete(id uint) errors.Errors {
	panic("unimplemented")
}

// Edit implements TreatmentService.
func (t *treatmentServiceImpl) Edit(id uint, payload *dto.TreatmentPayload) errors.Errors {
	panic("unimplemented")
}

// GetAll implements TreatmentService.
func (t *treatmentServiceImpl) GetAll() errors.Errors {
	panic("unimplemented")
}

// GetById implements TreatmentService.
func (t *treatmentServiceImpl) GetById(id uint) errors.Errors {
	panic("unimplemented")
}

func NewTreatmentServiceImpl(tr treatment_repository.TreatmentRepository, wg *sync.WaitGroup) TreatmentService {
	return &treatmentServiceImpl{
		tr: tr,
		wg: wg,
	}
}
