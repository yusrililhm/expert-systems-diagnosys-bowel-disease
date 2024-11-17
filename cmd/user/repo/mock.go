package repo

import (
	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/pkg/exception"
)

type userRepoMock struct {
}

var (
	Add            func(user *entity.User) exception.Exception
	FetchByPhone   func(phone string) (*entity.User, exception.Exception)
	ChangePassword func(user *entity.User) exception.Exception
	FetchById      func(id int) (*entity.User, exception.Exception)
	Modify         func(user *entity.User) exception.Exception
)

// Add implements UserRepo.
func (u *userRepoMock) Add(user *entity.User) exception.Exception {
	return Add(user)
}

// FetchByPhone implements UserRepo.
func (u *userRepoMock) FetchByPhone(phone string) (*entity.User, exception.Exception) {
	return FetchByPhone(phone)
}

// ChangePassword implements UserRepo.
func (u *userRepoMock) ChangePassword(user *entity.User) exception.Exception {
	return ChangePassword(user)
}

// FetchById implements UserRepo.
func (u *userRepoMock) FetchById(id int) (*entity.User, exception.Exception) {
	return FetchById(id)
}

// Modify implements UserRepo.
func (u *userRepoMock) Modify(user *entity.User) exception.Exception {
	return Modify(user)
}

func NewUserRepoMock() UserRepo {
	return &userRepoMock{}
}
