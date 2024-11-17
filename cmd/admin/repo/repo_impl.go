package repo

import (
	"errors"
	"log"

	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/exception"

	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

// FetchByUsername implements AdminRepo.
func (ar *adminRepo) FetchByUsername(username string) (*entity.User, exception.Exception) {
	u := &entity.User{}

	if err := ar.db.First(&u, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[warn] something happened here :", err.Error())
			return nil, exception.NewNotFound("Admin not found")
		}

		log.Println("[warn] something happened here :", err.Error())
		return nil, exception.NewInternalServerError("Something went wrong")
	}

	return u, nil
}

// FetchAllUser implements AdminRepo.
func (ar *adminRepo) FetchAllUser() ([]*model.User, exception.Exception) {
	users := []*model.User{}

	if err := ar.db.Model(&entity.User{}).Where("role = ?", "User").Find(&users).Error; err != nil {
		log.Println("[warn] something happened here :", err.Error())
		return nil, exception.NewInternalServerError("Something went wrong")
	}

	return users, nil
}

func NewAdminRepo(db *gorm.DB) AdminRepo {
	return &adminRepo{
		db: db,
	}
}
