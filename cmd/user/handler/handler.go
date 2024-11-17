package handler

import "net/http"

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Modify(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	ProfileView(w http.ResponseWriter, r *http.Request)
	RegisterView(w http.ResponseWriter, r *http.Request)
	LoginView(w http.ResponseWriter, r *http.Request)
	DiagnosysHistoryView(w http.ResponseWriter, r *http.Request)
	ChangePasswordView(w http.ResponseWriter, r *http.Request)
}
