package user_service

import (
	"healthy-bowel/internal/pkg/dto"
	"healthy-bowel/internal/pkg/errors"
)

type UserService interface {
	Add(payload *dto.AddUserPayload) errors.Errors
	Login(payload *dto.UserLoginPayload) errors.Errors
	Edit(id uint, payload *dto.EditUserPayload) errors.Errors
	Profile(id uint) errors.Errors
	ChangePassword(id uint, payload *dto.ChangePasswordPayload) errors.Errors
}
