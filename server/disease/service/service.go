package service

import (
	"usus-sehat/server/exception"
	"usus-sehat/server/model"
)

type DiseaseService interface {
	FetchAllDiseases() (*model.SuccessResponse, exception.Exception)
}
