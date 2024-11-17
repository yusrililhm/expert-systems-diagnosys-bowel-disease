package service

import (
	"fmt"
	"net/http"

	"usus-sehat/cmd/admin/repo"
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/dto"
	"usus-sehat/internal/pkg/exception"
)

type adminService struct {
	ar repo.AdminRepo
}

// FetchByUsername implements AdminService.
func (as *adminService) FetchByUsername(payload *dto.AdminLoginPayload) (*model.SuccessResponse, exception.Exception) {

	u, err := as.ar.FetchByUsername(payload.Username)

	if err != nil {
		return nil, err
	}

	if !u.CompareHashPassword(payload.Password) {
		return nil, exception.NewBadRequestError("Invalid admin")
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Admin successfully login",
		Data: &model.TokenResponse{
			Token: fmt.Sprintf("Bearer %s", u.GenerateTokenString()),
		},
	}, nil
}

// FindAllUsers implements AdminService.
func (as *adminService) FindAllUsers() (*model.SuccessResponse, exception.Exception) {

	users, err := as.ar.FetchAllUser()

	if err != nil {
		return nil, err
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Users successfully fetched",
		Data:    users,
	}, nil
}

func NewAdminService(ar repo.AdminRepo) AdminService {
	return &adminService{
		ar: ar,
	}
}
