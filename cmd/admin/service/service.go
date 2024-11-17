package service

import (
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/dto"
	"usus-sehat/internal/pkg/exception"
)

type AdminService interface {
	FindAllUsers() (*model.SuccessResponse, exception.Exception)
	FetchByUsername(payload *dto.AdminLoginPayload) (*model.SuccessResponse, exception.Exception)
}
