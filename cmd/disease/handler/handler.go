package handler

import "net/http"

type DiseaseHandler interface {
	DiseaseDashboardView(w http.ResponseWriter, r *http.Request)
}
