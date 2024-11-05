package service

import (
	"usus-sehat/server/exception"
	"usus-sehat/server/model"
)

type TreatmentService interface {
	FetchAllTreatments() (*model.SuccessResponse, exception.Exception)
}
