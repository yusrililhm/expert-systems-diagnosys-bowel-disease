package service

import (
	"fmt"
	"net/http"

	"usus-sehat/cmd/user/repo"
	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/dto"
	"usus-sehat/internal/pkg/exception"

	"gorm.io/gorm"
)

type userService struct {
	ur repo.UserRepo
}

// ChangePassword implements UserService.
func (us *userService) ChangePassword(id int, payload *dto.ChangePasswordPayload) (*model.SuccessResponse, exception.Exception) {

	u, err := us.ur.FetchById(id)

	if err != nil {
		return nil, err
	}

	if !u.CompareHashPassword(payload.OldPassword) {
		return nil, exception.NewBadRequestError("Invalid old password")
	}

	if payload.NewPassword != payload.ConfirmNewPassword {
		return nil, exception.NewBadRequestError("Invalid confirm new password")
	}

	user := &entity.User{
		Model:    gorm.Model{ID: uint(id)},
		Password: payload.NewPassword,
	}

	user.GenerateFromPassword()

	if err := us.ur.ChangePassword(user); err != nil {
		return nil, err
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Password successfully changed",
		Data:    nil,
	}, nil
}

// Modify implements UserService.
func (us *userService) Modify(id int, payload *dto.UserModifyPayload) (*model.SuccessResponse, exception.Exception) {

	_, err := us.ur.FetchById(id)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Model:     gorm.Model{ID: uint(id)},
		FullName:  payload.FullName,
		Phone:     payload.Phone,
		BirthDate: payload.BirthDate,
	}

	if err := us.ur.Modify(user); err != nil {
		return nil, err
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "User successfully modified",
		Data:    nil,
	}, nil
}

// Profile implements UserService.
func (us *userService) Profile(id int) (*model.SuccessResponse, exception.Exception) {

	u, err := us.ur.FetchById(id)

	if err != nil {
		return nil, err
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "",
		Data:    u,
	}, nil
}

// Login implements UserService.
func (us *userService) Login(payload *dto.UserLoginPayload) (*model.SuccessResponse, exception.Exception) {

	u, err := us.ur.FetchByPhone(payload.Phone)

	if err != nil {
		return nil, err
	}

	if !u.CompareHashPassword(payload.Password) {
		return nil, exception.NewBadRequestError("Invalid phone or password")
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "User successfully login",
		Data: &model.TokenResponse{
			Token: fmt.Sprintf("Bearer %s", u.GenerateTokenString()),
		},
	}, nil
}

// Register implements UserService.
func (us *userService) Register(payload *dto.UserRegisterPayload) (*model.SuccessResponse, exception.Exception) {

	user := &entity.User{
		Username:  payload.Username,
		FullName:  payload.FullName,
		Phone:     payload.Phone,
		BirthDate: payload.BirthDate,
		Role:      "User",
		Gender:    payload.Gender,
		Password:  payload.Password,
	}

	user.GenerateFromPassword()

	if err := us.ur.Add(user); err != nil {
		return nil, err
	}

	return &model.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "User successfully registered",
		Data:    nil,
	}, nil
}

func NewUserService(ur repo.UserRepo) UserService {
	return &userService{
		ur: ur,
	}
}
