package service

import (
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/dto"
	"usus-sehat/internal/pkg/exception"
)

type UserService interface {
	Register(payload *dto.UserRegisterPayload) (*model.SuccessResponse, exception.Exception)
	Login(payload *dto.UserLoginPayload) (*model.SuccessResponse, exception.Exception)
	Profile(id int) (*model.SuccessResponse, exception.Exception)
	Modify(id int, payload *dto.UserModifyPayload) (*model.SuccessResponse, exception.Exception)
	ChangePassword(id int, payload *dto.ChangePasswordPayload) (*model.SuccessResponse, exception.Exception)
}
