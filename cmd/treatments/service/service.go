package service

import (
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/exception"
)

type TreatmentService interface {
	FetchAllTreatments() (*model.SuccessResponse, exception.Exception)
}
