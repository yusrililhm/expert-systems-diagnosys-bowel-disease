package service

import (
	"net/http"

	"usus-sehat/cmd/treatments/repo"
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/exception"
)

type treatmentService struct {
	tr repo.TreatmentRepo
}

// FetchAllTreatments implements TreatmentService.
func (ts *treatmentService) FetchAllTreatments() (*model.SuccessResponse, exception.Exception) {
	treatments, err := ts.tr.FetchAllTreatments()

	if err != nil {
		return nil, err
	}

	return &model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Treatments successfully fetched",
		Data:    treatments,
	}, nil
}

func NewTreatmentService(tr repo.TreatmentRepo) TreatmentService {
	return &treatmentService{
		tr: tr,
	}
}
