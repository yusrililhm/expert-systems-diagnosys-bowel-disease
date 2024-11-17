package handler

import "net/http"

type SymptomHandler interface {
	SymptomDashboardView(w http.ResponseWriter, r *http.Request)
}
