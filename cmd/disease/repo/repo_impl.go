package repo

import (
	"log"

	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/pkg/exception"

	"gorm.io/gorm"
)

type diseaseRepo struct {
	db *gorm.DB
}

// FetchAllDiseases implements DiseaseRepo.
func (dr *diseaseRepo) FetchAllDiseases() ([]*entity.Disease, exception.Exception) {
	diseases := []*entity.Disease{}

	if err := dr.db.Model(&entity.Disease{}).Preload("Symptoms").Preload("Treatments").Find(&diseases).Error; err != nil {
		log.Println("[warn] An error occured :", err.Error())
		return nil, exception.NewInternalServerError("Something went wrong")
	}

	return diseases, nil
}

func NewDiseaseRepo(db *gorm.DB) DiseaseRepo {
	return &diseaseRepo{
		db: db,
	}
}
