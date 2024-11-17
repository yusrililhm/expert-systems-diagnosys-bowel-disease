package handler

import "net/http"

type TreatmentHandler interface {
	TreatmentsDashboardView(w http.ResponseWriter, r *http.Request)
}
