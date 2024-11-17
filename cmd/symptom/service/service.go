package service

import (
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/exception"
)

type SymptomService interface {
	FindAllSymptoms() (*model.SuccessResponse, exception.Exception)
}
