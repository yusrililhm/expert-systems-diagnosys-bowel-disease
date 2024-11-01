package service

import (
	"net/http"
	"usus-sehat/server/exception"
	"usus-sehat/server/model"
	"usus-sehat/server/symptom/repo"
)

type symptomService struct {
	sr repo.SymptomRepo
}

// FindAllSymptoms implements SymptomService.
func (ss *symptomService) FindAllSymptoms() (*model.SuccessResponse, exception.Exception) {

	symptoms, err := ss.sr.FindAllSymptoms()

	if err != nil {
		return nil, err
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Symptoms successfull fetched",
		Data:    symptoms,
	}, nil
}

func NewSymptomService(sr repo.SymptomRepo) SymptomService {
	return &symptomService{
		sr: sr,
	}
}
