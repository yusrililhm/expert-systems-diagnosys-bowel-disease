package handler

import (
	"html/template"
	"log"
	"net/http"
	"usus-sehat/server/disease/service"
	"usus-sehat/server/model"
)

type diseaseHandler struct {
	ds service.DiseaseService
}

// DiseaseDashboardView implements DiseaseHandler.
func (dh *diseaseHandler) DiseaseDashboardView(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(
		"web/template/views/admin/disease.html",
		model.Header,
		model.AdminNavbar,
	)

	if err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	res, e := dh.ds.FetchAllDiseases()

	if e != nil {
		w.WriteHeader(e.Status())
		w.Write(model.ToJSON(e))

		return
	}

	if err := templ.ExecuteTemplate(w, "disease", res); err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func NewDiseaseHandler(ds service.DiseaseService) DiseaseHandler {
	return &diseaseHandler{
		ds: ds,
	}
}
