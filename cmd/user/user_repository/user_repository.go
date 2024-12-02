package user_repository

import (
	"healthy-bowel/internal/domain"
	"healthy-bowel/internal/pkg/errors"
)

type UserRepository interface {
	Add(user *domain.User) errors.Errors
	GetById(id uint) (*domain.User, errors.Errors)
	GetByUsername(username string) (*domain.User, errors.Errors)
	GetAll() ([]*domain.User, errors.Errors)
	Edit(user *domain.User) errors.Errors
	ChanngePassword(user *domain.User) errors.Errors
}
