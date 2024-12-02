package user_service

import (
	"healthy-bowel/cmd/user/user_repository"
	"healthy-bowel/internal/pkg/dto"
	"healthy-bowel/internal/pkg/errors"
	"sync"
)

type userServiceImpl struct {
	ur user_repository.UserRepository
	wg *sync.WaitGroup
}

// Add implements UserService.
func (us *userServiceImpl) Add(payload *dto.AddUserPayload) errors.Errors {
	panic("unimplemented")
}

// ChangePassword implements UserService.
func (us *userServiceImpl) ChangePassword(id uint, payload *dto.ChangePasswordPayload) errors.Errors {
	panic("unimplemented")
}

// Edit implements UserService.
func (us *userServiceImpl) Edit(id uint, payload *dto.EditUserPayload) errors.Errors {
	panic("unimplemented")
}

// Login implements UserService.
func (us *userServiceImpl) Login(payload *dto.UserLoginPayload) errors.Errors {
	panic("unimplemented")
}

// Profile implements UserService.
func (us *userServiceImpl) Profile(id uint) errors.Errors {
	panic("unimplemented")
}

func NewUserServiceImpl(ur user_repository.UserRepository, wg *sync.WaitGroup) UserService {
	return &userServiceImpl{
		ur: ur,
		wg: wg,
	}
}
