package service

import (
	"net/http"

	"usus-sehat/cmd/disease/repo"
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/exception"
)

type diseaseService struct {
	dr repo.DiseaseRepo
}

// FetchAllDiseases implements DiseaseService.
func (ds *diseaseService) FetchAllDiseases() (*model.SuccessResponse, exception.Exception) {

	diseases, err := ds.dr.FetchAllDiseases()

	if err != nil {
		return nil, err
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Diseases successfully fetched",
		Data:    diseases,
	}, nil
}

func NewDiseaseService(dr repo.DiseaseRepo) DiseaseService {
	return &diseaseService{
		dr: dr,
	}
}
