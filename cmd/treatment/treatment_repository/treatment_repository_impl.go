package treatment_repository

import (
	"database/sql"
	"healthy-bowel/internal/domain"
	"healthy-bowel/internal/pkg/errors"
	"sync"
)

type treatmentRepositoryImpl struct {
	db *sql.DB
	wg *sync.WaitGroup
}

// Add implements TreatmentRepository.
func (tr *treatmentRepositoryImpl) Add(treatment *domain.Treatment) errors.Errors {
	panic("unimplemented")
}

// Delete implements TreatmentRepository.
func (tr *treatmentRepositoryImpl) Delete(id uint) errors.Errors {
	panic("unimplemented")
}

// Edit implements TreatmentRepository.
func (tr *treatmentRepositoryImpl) Edit(treatment *domain.Treatment) errors.Errors {
	panic("unimplemented")
}

// GetAll implements TreatmentRepository.
func (tr *treatmentRepositoryImpl) GetAll() ([]*domain.Treatment, errors.Errors) {
	panic("unimplemented")
}

// GetById implements TreatmentRepository.
func (tr *treatmentRepositoryImpl) GetById(id uint) (*domain.Treatment, errors.Errors) {
	panic("unimplemented")
}

func NewTreatmentRepositoryImpl(db *sql.DB, wg *sync.WaitGroup) TreatmentRepository {
	return &treatmentRepositoryImpl{
		db: db,
		wg: wg,
	}
}
