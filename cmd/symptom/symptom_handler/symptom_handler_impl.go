package symptom_handler

import (
	"sync"

	"healthy-bowel/cmd/symptom/symptom_service"

	"github.com/gofiber/fiber/v2"
)

type symptomHandlerImpl struct {
	ss symptom_service.SymptomService
	wg *sync.WaitGroup
}

// Add implements SymptomHandler.
func (sh *symptomHandlerImpl) Add(c *fiber.Ctx) error {
	return c.Render("", "")
}

// Delete implements SymptomHandler.
func (sh *symptomHandlerImpl) Delete(c *fiber.Ctx) error {
	return c.Render("", "")
}

// Edit implements SymptomHandler.
func (sh *symptomHandlerImpl) Edit(c *fiber.Ctx) error {
	return c.Render("", "")
}

// GetAll implements SymptomHandler.
func (sh *symptomHandlerImpl) GetAll(c *fiber.Ctx) error {
	return c.Render("", "")
}

// GetById implements SymptomHandler.
func (sh *symptomHandlerImpl) GetById(c *fiber.Ctx) error {
	return c.Render("", "")
}

func NewSymptomHandlerImpl(ss symptom_service.SymptomService, wg *sync.WaitGroup) SymptomHandler {
	return &symptomHandlerImpl{
		ss: ss,
		wg: wg,
	}
}
