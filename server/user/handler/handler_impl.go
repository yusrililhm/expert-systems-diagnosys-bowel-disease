package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"usus-sehat/server/dto"
	"usus-sehat/server/entity"
	"usus-sehat/server/helper"
	"usus-sehat/server/middleware"
	"usus-sehat/server/model"
	"usus-sehat/server/user/service"

	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
)

type userHandler struct {
	us service.UserService
	sc *securecookie.SecureCookie
}

// ChangePasswordView implements UserHandler.
func (uh *userHandler) ChangePasswordView(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(helper.Header, helper.UserNavbar, "web/template/views/users/change_password.html")

	if err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.ExecuteTemplate(w, "change_password", map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r)}); err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DiagnosysHistoryView implements UserHandler.
func (uh *userHandler) DiagnosysHistoryView(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(helper.Header, helper.UserNavbar, "web/template/views/users/diagnosys_history.html")

	if err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.ExecuteTemplate(w, "diagnosys_history", nil); err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// LoginView implements UserHandler.
func (uh *userHandler) LoginView(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(helper.Header, "web/template/views/users/login.html")

	if err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.ExecuteTemplate(w, "login", map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r)}); err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RegisterView implements UserHandler.
func (uh *userHandler) RegisterView(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(helper.Header, "web/template/views/users/register.html")

	if err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.ExecuteTemplate(w, "register", map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r)}); err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ChangePassword implements UserHandler.
func (uh *userHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value(middleware.Key).(*entity.User).ID

	if err := r.ParseForm(); err != nil {
		log.Println("[warn] An error occured :", err.Error())
		http.Error(w, "Sommething went wrong", http.StatusInternalServerError)
		return
	}

	payload := &dto.ChangePasswordPayload{
		OldPassword:        r.FormValue("old_password"),
		NewPassword:        r.FormValue("new_password"),
		ConfirmNewPassword: r.FormValue("confirm_new_password"),
	}

	if err := helper.ValidatePayload(payload); err != nil {
		http.Error(w, err.Message(), err.Status())
		return
	}

	_, err := uh.us.ChangePassword(int(id), payload)

	if err != nil {
		http.Error(w, err.Message(), err.Status())
		return
	}

	http.Redirect(w, r, "/change-password", http.StatusFound)
}

// Login implements UserHandler.
func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println("[warn] An error occured :", err.Error())
		http.Error(w, "Sommething went wrong", http.StatusInternalServerError)
		return
	}

	payload := &dto.UserLoginPayload{
		Phone:    r.FormValue("phone"),
		Password: r.FormValue("password"),
	}

	if err := helper.ValidatePayload(payload); err != nil {
		http.Error(w, err.Message(), err.Status())
		return
	}

	res, err := uh.us.Login(payload)

	if err != nil {
		http.Error(w, err.Message(), err.Status())
		return
	}

	encoded, e := uh.sc.Encode("token", res.Data.(*model.TokenResponse))

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
	http.Redirect(w, r, "/profile", http.StatusFound)
}

// Modify implements UserHandler.
func (uh *userHandler) Modify(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.Key).(*entity.User).ID

	if err := r.ParseForm(); err != nil {
		log.Println("[warn] An error occured :", err.Error())
		http.Error(w, "Sommething went wrong", http.StatusInternalServerError)
		return
	}

	date, err := time.Parse("2006-01-02", r.FormValue("birth_date"))

	if err != nil {
		log.Println("[warn] An error occured :", err.Error())
		http.Error(w, "Sommething went wrong", http.StatusInternalServerError)
		return
	}

	payload := &dto.UserModifyPayload{
		Username:  r.FormValue("username"),
		FullName:  r.FormValue("full_name"),
		Phone:     r.FormValue("phone"),
		BirthDate: date,
	}

	_, e := uh.us.Modify(int(id), payload)

	if e != nil {
		http.Error(w, e.Message(), e.Status())
		return
	}

	http.Redirect(w, r, "/profile", http.StatusFound)
}

// ProfileView implements UserHandler.
func (uh *userHandler) ProfileView(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value(middleware.Key).(*entity.User).ID

	res, e := uh.us.Profile(int(id))

	if e != nil {
		http.Error(w, e.Message(), e.Status())
		return
	}

	temp, err := template.ParseFiles(helper.Header, helper.UserNavbar, "web/template/views/users/profile.html")

	if err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.ExecuteTemplate(w, "profile", res); err != nil {
		log.Println("[warn]", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Register implements UserHandler.
func (uh *userHandler) Register(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println("[warn] An error occured :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	date, err := time.Parse("2006-01-02", r.FormValue("birth_date"))

	if err != nil {
		log.Println("[warn] An error occured :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	gender, err := strconv.ParseBool(r.FormValue("gender"))
	if err != nil {
		log.Println("[warn] An error occured :", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	payload := &dto.UserRegisterPayload{
		Username:  r.FormValue("username"),
		FullName:  r.FormValue("full_name"),
		Phone:     r.FormValue("phone"),
		BirthDate: date,
		Gender:    gender,
		Password:  r.FormValue("password"),
	}

	if err := helper.ValidatePayload(payload); err != nil {
		http.Error(w, err.Message(), err.Status())
		return
	}

	_, e := uh.us.Register(payload)

	if e != nil {
		http.Error(w, e.Message(), e.Status())
		return
	}

	http.Redirect(w, r, "/auth/login", http.StatusFound)
}

func NewUserHandler(us service.UserService, sc *securecookie.SecureCookie) UserHandler {
	return &userHandler{
		us: us,
		sc: sc,
	}
}
