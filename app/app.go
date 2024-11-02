package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	adminHandler "usus-sehat/server/admin/handler"
	adminRepo "usus-sehat/server/admin/repo"
	adminService "usus-sehat/server/admin/service"

	"usus-sehat/server/middleware"

	symptomHandler "usus-sehat/server/symptom/handler"
	symptomRepo "usus-sehat/server/symptom/repo"
	symptomService "usus-sehat/server/symptom/service"

	diseaseHandler "usus-sehat/server/disease/handler"
	diseaseRepo "usus-sehat/server/disease/repo"
	diseaseService "usus-sehat/server/disease/service"

	"usus-sehat/server/config"
	"usus-sehat/server/infra/db"

	userHandler "usus-sehat/server/user/handler"
	userRepo "usus-sehat/server/user/repo"
	userService "usus-sehat/server/user/service"

	"github.com/gorilla/securecookie"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func StartNonTLSServer() {
	mux := chi.NewMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://localhost", http.StatusTemporaryRedirect)
	}))

	http.ListenAndServe(":80", mux)
}

func StartApp() {

	// load env file
	config.LoadEnv()

	// connect to db
	db, err := db.NewDb()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// secure cookie
	sc := securecookie.New(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(16))

	// dependency injection
	ur := userRepo.NewUserRepo(db)
	us := userService.NewUserService(ur)
	uh := userHandler.NewUserHandler(us, sc)

	ar := adminRepo.NewAdminRepo(db)
	as := adminService.NewAdminService(ar)
	ah := adminHandler.NewAdminHandler(as, sc)

	sr := symptomRepo.NewSymptomRepo(db)
	ss := symptomService.NewSymptomService(sr)
	sh := symptomHandler.NewSymptomHandler(ss)

	dr := diseaseRepo.NewDiseaseRepo(db)
	ds := diseaseService.NewDiseaseService(dr)
	dh := diseaseHandler.NewDiseaseHandler(ds)

	md := middleware.NewMiddleware(sc, ur)

	// router
	r := chi.NewRouter()

	// csrf
	// r.Use(csrf.Protect(securecookie.GenerateRandomKey(32), csrf.Secure(false)))

	// logger
	r.Use(chiMiddleware.Logger)

	// recoverer
	r.Use(chiMiddleware.Recoverer)

	// user group routes
	r.Group(func(r chi.Router) {
		r.Post("/api/v1/register", uh.Register)
		r.Post("/api/v1/login", uh.Login)

		r.Get("/auth/register", uh.RegisterView)
		r.Get("/auth/login", uh.LoginView)

		r.Group(func(r chi.Router) {
			r.Use(md.Authentication)

			r.Patch("/api/v1/change-password", uh.ChangePassword)
			r.Patch("/api/v1/modify", uh.Modify)

			r.Get("/profile", uh.ProfileView)
			r.Get("/change-password", uh.ChangePasswordView)
			r.Get("/diagnosys-history", uh.DiagnosysHistoryView)
		})
	})

	// admin group routes
	r.Group(func(r chi.Router) {
		r.Get("/admin/login", ah.AdminLoginView)
		r.Post("/api/v1/admin/login", ah.AdminLogin)

		r.Group(func(r chi.Router) {
			r.Use(md.Authentication, md.Authorization)
			r.Get("/admin/dashboard", ah.DashboardView)
			r.Get("/admin/diseases", dh.DiseaseDashboardView)
			r.Get("/admin/symptoms", sh.SymptomDashboardView)
			r.Get("/admin/treatments", ah.DashboardView)
		})
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/template/assets/"))))

	log.Printf("Server is running on PORT %s 🚀\n", os.Getenv("APP_PORT"))

	if err := http.ListenAndServeTLS(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), "server.crt", "server.key", r); err != nil {
		log.Fatal(err.Error())
		return
	}
}
