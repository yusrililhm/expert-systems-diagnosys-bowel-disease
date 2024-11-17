package handler

import "net/http"

type AdminHandler interface {
	DashboardView(w http.ResponseWriter, r *http.Request)
	AdminLoginView(w http.ResponseWriter, r *http.Request)
	AdminLogin(w http.ResponseWriter, r *http.Request)
}
