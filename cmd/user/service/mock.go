package service

import (
	"net/http"
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/dto"
	"usus-sehat/internal/pkg/exception"
)

type userServiceMock struct {
}

var (
	Authentication func(next http.Handler) http.Handler
	Authorization  func(next http.Handler) http.Handler
	Login          func(payload *dto.UserLoginPayload) (*model.SuccessResponse, exception.Exception)
	Register       func(payload *dto.UserRegisterPayload) (*model.SuccessResponse, exception.Exception)
	ChangePassword func(id int, payload *dto.ChangePasswordPayload) (*model.SuccessResponse, exception.Exception)
	Modify         func(id int, payload *dto.UserModifyPayload) (*model.SuccessResponse, exception.Exception)
	Profile        func(id int) (*model.SuccessResponse, exception.Exception)
)

// Authentication implements UserService.
func (u *userServiceMock) Authentication(next http.Handler) http.Handler {
	return Authentication(next)
}

// Authorization implements UserService.
func (u *userServiceMock) Authorization(next http.Handler) http.Handler {
	return Authorization(next)
}

// Login implements UserService.
func (u *userServiceMock) Login(payload *dto.UserLoginPayload) (*model.SuccessResponse, exception.Exception) {
	return Login(payload)
}

// Register implements UserService.
func (u *userServiceMock) Register(payload *dto.UserRegisterPayload) (*model.SuccessResponse, exception.Exception) {
	return Register(payload)
}

// ChangePassword implements UserService.
func (u *userServiceMock) ChangePassword(id int, payload *dto.ChangePasswordPayload) (*model.SuccessResponse, exception.Exception) {
	return ChangePassword(id, payload)
}

// Modify implements UserService.
func (u *userServiceMock) Modify(id int, payload *dto.UserModifyPayload) (*model.SuccessResponse, exception.Exception) {
	return Modify(id, payload)
}

// Profile implements UserService.
func (u *userServiceMock) Profile(id int) (*model.SuccessResponse, exception.Exception) {
	return Profile(id)
}

func NewUserServiceMock() UserService {
	return &userServiceMock{}
}
