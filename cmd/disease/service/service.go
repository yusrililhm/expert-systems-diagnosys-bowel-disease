package service

import (
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/exception"
)

type DiseaseService interface {
	FetchAllDiseases() (*model.SuccessResponse, exception.Exception)
}
