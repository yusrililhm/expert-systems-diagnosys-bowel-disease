package handler

import (
	"log"
	"net/http"
	"text/template"
	"time"
	"usus-sehat/server/admin/service"
	"usus-sehat/server/dto"
	"usus-sehat/server/model"

	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
)

type adminHandler struct {
	as service.AdminService
	sc *securecookie.SecureCookie
}

// AdminLogin implements AdminHandler.
func (ah *adminHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println("[warn] something happened here :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	payload := &dto.AdminLoginPayload{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	res, err := ah.as.FetchByUsername(payload)

	if err != nil {
		http.Error(w, err.Message(), err.Status())
		return
	}

	encoded, e := ah.sc.Encode("token", res.Data.(*model.TokenResponse))

	if e != nil {
		log.Println("[warn]", e.Error())
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   encoded,
		Path:    "/",
		Expires: time.Now().Add(2 * time.Minute),
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
}

// AdminLoginView implements AdminHandler.
func (ah *adminHandler) AdminLoginView(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(
		"web/template/views/admin/login.html",
		model.Header,
	)

	if err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err := templ.ExecuteTemplate(w, "admin_login", map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r)}); err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

// DashboardView implements AdminHandler.
func (ah *adminHandler) DashboardView(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(
		"web/template/views/admin/dashboard.html",
		model.Header,
		model.AdminNavbar,
	)

	if err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	res, e := ah.as.FindAllUsers()

	if e != nil {
		w.WriteHeader(e.Status())
		w.Write(model.ToJSON(e))

		return
	}

	if err := templ.ExecuteTemplate(w, "dashboard", res); err != nil {
		log.Println("[warn] An error occured when rendering template :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func NewAdminHandler(as service.AdminService, sc *securecookie.SecureCookie) AdminHandler {
	return &adminHandler{
		as: as,
		sc: sc,
	}
}
