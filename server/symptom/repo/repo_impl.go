package repo

import (
	"log"
	"usus-sehat/server/entity"
	"usus-sehat/server/exception"

	"gorm.io/gorm"
)

type symptomRepo struct {
	db *gorm.DB
}

// FindAllSymptoms implements SymptomRepo.
func (sr *symptomRepo) FindAllSymptoms() ([]*entity.Symptom, exception.Exception) {
	symptoms := []*entity.Symptom{}

	if err := sr.db.Model(&entity.Symptom{}).Find(&symptoms).Error; err != nil {
		log.Println("[warn] An error occured :", err.Error())
		return nil, exception.NewInternalServerError("Something went wrong")
	}

	return symptoms, nil
}

func NewSymptomRepo(db *gorm.DB) SymptomRepo {
	return &symptomRepo{
		db: db,
	}
}
