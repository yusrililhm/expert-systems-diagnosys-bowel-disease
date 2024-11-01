package service

import (
	"net/http"
	"usus-sehat/server/disease/repo"
	"usus-sehat/server/exception"
	"usus-sehat/server/model"
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
