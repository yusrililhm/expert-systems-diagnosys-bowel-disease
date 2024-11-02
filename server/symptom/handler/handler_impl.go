package handler

import (
	"html/template"
	"log"
	"net/http"
	"usus-sehat/server/helper"
	"usus-sehat/server/symptom/service"
)

type symptomHandler struct {
	ss service.SymptomService
}

// SymptomDashboardView implements SymptomHandler.
func (sh *symptomHandler) SymptomDashboardView(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Encoding", "gzip")

	templ, err := template.ParseFiles(
		"web/template/views/admin/symptoms.html",
		helper.Header,
		helper.AdminNavbar,
	)

	if err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	res, e := sh.ss.FindAllSymptoms()

	if e != nil {
		w.WriteHeader(e.Status())
		w.Write(helper.ToJSON(e))

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
