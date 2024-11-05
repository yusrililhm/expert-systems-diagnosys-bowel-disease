package handler

import (
	"html/template"
	"log"
	"net/http"
	"usus-sehat/server/helper"
	"usus-sehat/server/treatments/service"
)

type treatmentHandler struct {
	ts service.TreatmentService
}

// TreatmentsDashboardView implements TreatmentHandler.
func (th *treatmentHandler) TreatmentsDashboardView(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(
		"web/template/views/admin/treatments.html",
		helper.Header,
		helper.AdminNavbar,
	)

	if err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	res, e := th.ts.FetchAllTreatments()

	if e != nil {
		w.WriteHeader(e.Status())
		w.Write(helper.ToJSON(e))

		return
	}

	if err := templ.ExecuteTemplate(w, "treatment", res); err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func NewTreatmentHandler(ts service.TreatmentService) TreatmentHandler {
	return &treatmentHandler{
		ts: ts,
	}
}
