package handler

import (
	"html/template"
	"log"
	"net/http"

	"usus-sehat/cmd/symptom/service"
	"usus-sehat/internal/domain/model"
)

type symptomHandler struct {
	ss service.SymptomService
}

// SymptomDashboardView implements SymptomHandler.
func (sh *symptomHandler) SymptomDashboardView(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(
		"web/template/views/admin/symptoms.html",
		model.Header,
		model.AdminNavbar,
	)

	if err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	res, e := sh.ss.FindAllSymptoms()

	if e != nil {
		w.WriteHeader(e.Status())
		w.Write(model.ToJSON(e))

		return
	}

	if err := templ.ExecuteTemplate(w, "symptom", res); err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func NewSymptomHandler(ss service.SymptomService) SymptomHandler {
	return &symptomHandler{
		ss: ss,
	}
}
