package app

import (
	"log"
	"os"
	"sync"

	"healthy-bowel/cmd/symptom/symptom_handler"
	"healthy-bowel/cmd/symptom/symptom_repository"
	"healthy-bowel/cmd/symptom/symptom_service"

	"healthy-bowel/cmd/treatment/treatment_handler"
	"healthy-bowel/cmd/treatment/treatment_repository"
	"healthy-bowel/cmd/treatment/treatment_service"

	"healthy-bowel/cmd/user/user_handler"
	"healthy-bowel/cmd/user/user_repository"
	"healthy-bowel/cmd/user/user_service"

	"healthy-bowel/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewApp() {
	db, err := db.NewDb()

	if err != nil {
		log.Fatal(err.Error())
	}

	wg := &sync.WaitGroup{}

	ur := user_repository.NewUserRepositoryImpl(db, wg)
	us := user_service.NewUserServiceImpl(ur, wg)
	uh := user_handler.NewUserHandlerImpl(us, wg)

	tr := treatment_repository.NewTreatmentRepositoryImpl(db, wg)
	ts := treatment_service.NewTreatmentServiceImpl(tr, wg)
	th := treatment_handler.NewTreatmentHandlerImpl(ts, wg)

	sr := symptom_repository.NewSymptomRepositoryImpl(db, wg)
	ss := symptom_service.NewSymptomServiceImpl(sr, wg)
	sh := symptom_handler.NewSymptomHandlerImpl(ss, wg)

	r := fiber.New(fiber.Config{
		Views: html.New("web/templates", ".html"),
	})

	users := r.Group("/user")

	users.Post("/register", uh.Register)
	users.Post("/login", uh.Login)
	users.Get("/", uh.Profile)
	users.Patch("/edit", uh.Edit)
	users.Patch("/change-password", uh.ChangePassword)

	treatments := r.Group("/treatments")

	treatments.Post("/", th.Add)
	treatments.Get("/", th.GetAll)
	treatments.Get("/:id", th.GetById)
	treatments.Patch("/:id", th.Edit)
	treatments.Delete("/:id", th.Delete)

	symptoms := r.Group("/symptoms")

	symptoms.Post("/", sh.Add)
	symptoms.Get("/", sh.GetAll)
	symptoms.Get("/:id", sh.GetById)
	symptoms.Patch("/:id", sh.Edit)
	symptoms.Delete("/:id", sh.Delete)

	r.Static("/static/", "web/static/")

	if err := r.ListenTLS(":"+os.Getenv("APP_PORT"), "sercer.crt", "server.key"); err != nil {
		log.Fatal(err.Error())
		return
	}
}
