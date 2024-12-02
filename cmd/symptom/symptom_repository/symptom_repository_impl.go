package symptom_repository

import (
	"database/sql"
	"healthy-bowel/internal/domain"
	"healthy-bowel/internal/pkg/errors"
	"sync"
)

type symptomRepositoryImpl struct {
	db *sql.DB
	wg *sync.WaitGroup
}

// Add implements SymptomRepository.
func (sr *symptomRepositoryImpl) Add(symptom *domain.Symptom) errors.Errors {
	panic("unimplemented")
}

// Delete implements SymptomRepository.
func (sr *symptomRepositoryImpl) Delete(id uint) errors.Errors {
	panic("unimplemented")
}

// Edit implements SymptomRepository.
func (sr *symptomRepositoryImpl) Edit(symptom *domain.Symptom) errors.Errors {
	panic("unimplemented")
}

// GetAll implements SymptomRepository.
func (sr *symptomRepositoryImpl) GetAll() ([]*domain.Symptom, errors.Errors) {
	panic("unimplemented")
}

// GetById implements SymptomRepository.
func (sr *symptomRepositoryImpl) GetById(id uint) (*domain.Symptom, errors.Errors) {
	panic("unimplemented")
}

func NewSymptomRepositoryImpl(db *sql.DB, wg *sync.WaitGroup) SymptomRepository {
	return &symptomRepositoryImpl{
		db: db,
		wg: wg,
	}
}
