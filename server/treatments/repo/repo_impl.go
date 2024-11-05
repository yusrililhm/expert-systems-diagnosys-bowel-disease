package repo

import (
	"log"
	"usus-sehat/server/entity"
	"usus-sehat/server/exception"

	"gorm.io/gorm"
)

type treatmentRepo struct {
	db *gorm.DB
}

// FetchAllTreatments implements TreatmentRepo.
func (tr *treatmentRepo) FetchAllTreatments() ([]*entity.Treatment, exception.Exception) {
	treatments := []*entity.Treatment{}

	if err := tr.db.Model(&entity.Treatment{}).Find(&treatments).Error; err != nil {
		log.Println("[warn] an error occured", err.Error())
		return nil, exception.NewInternalServerError("Something went wrong")
	}

	return treatments, nil
}

func NewTreatmentRepo(db *gorm.DB) TreatmentRepo {
	return &treatmentRepo{
		db: db,
	}
}
