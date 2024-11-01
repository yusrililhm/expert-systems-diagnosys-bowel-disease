package service

import (
	"usus-sehat/server/dto"
	"usus-sehat/server/exception"
	"usus-sehat/server/model"
)

type AdminService interface {
	FindAllUsers() (*model.SuccessResponse, exception.Exception)
	FetchByUsername(payload *dto.AdminLoginPayload) (*model.SuccessResponse, exception.Exception)
}
