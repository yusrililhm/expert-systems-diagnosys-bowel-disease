package repo

import (
	"errors"
	"log"
	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/pkg/exception"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// ChangePassword implements UserRepo.
func (ur *userRepo) ChangePassword(user *entity.User) exception.Exception {

	if err := ur.db.Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		log.Println("[warn]", err.Error())
		return exception.NewInternalServerError("Something went wrong")
	}

	return nil
}

// FetchById implements UserRepo.
func (ur *userRepo) FetchById(id int) (*entity.User, exception.Exception) {
	u := &entity.User{}

	if err := ur.db.First(&u, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[warn]", err.Error())
			return nil, exception.NewNotFound("User not found")
		}

		log.Println("[wanr]", err.Error())
		return nil, exception.NewInternalServerError("Something went wrong")
	}

	return u, nil
}

// Modify implements UserRepo.
func (ur *userRepo) Modify(user *entity.User) exception.Exception {

	if err := ur.db.Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		log.Println("[warn]", err.Error())
		return exception.NewInternalServerError("Something went wrong")
	}

	return nil
}

// FetchByPhone implements UserRepo.
func (ur *userRepo) FetchByPhone(phone string) (*entity.User, exception.Exception) {
	u := &entity.User{}

	if err := ur.db.First(&u, "phone = ?", phone).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[warn]", err.Error())
			return nil, exception.NewNotFound("User not found")
		}

		log.Println("[wanr]", err.Error())
		return nil, exception.NewInternalServerError("Something went wrong")
	}

	return u, nil
}

// Add implements UserRepo.
func (ur *userRepo) Add(user *entity.User) exception.Exception {

	if err := ur.db.Create(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("[warn] something happend here : %s", err.Error())
			return exception.NewConflictError("Username or phone has been used")
		}

		log.Printf("[warn] something happend here : %s", err.Error())
		return exception.NewInternalServerError("Something went wrong")
	}

	return nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
