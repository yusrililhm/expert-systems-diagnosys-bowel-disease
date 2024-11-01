package service

import (
	"usus-sehat/server/exception"
	"usus-sehat/server/model"
)

type SymptomService interface {
	FindAllSymptoms() (*model.SuccessResponse, exception.Exception)
}
