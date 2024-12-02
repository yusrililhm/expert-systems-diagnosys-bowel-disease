package treatment_handler

import (
	"healthy-bowel/cmd/treatment/treatment_service"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type treatmentHandlerImpl struct {
	ts treatment_service.TreatmentService
	wg *sync.WaitGroup
}

// Add implements TreatmentHandler.
func (t *treatmentHandlerImpl) Add(c *fiber.Ctx) error {
	return c.Render("", "")
}

// Delete implements TreatmentHandler.
func (t *treatmentHandlerImpl) Delete(c *fiber.Ctx) error {
	return c.Render("", "")
}

// Edit implements TreatmentHandler.
func (t *treatmentHandlerImpl) Edit(c *fiber.Ctx) error {
	return c.Render("", "")
}

// GetAll implements TreatmentHandler.
func (t *treatmentHandlerImpl) GetAll(c *fiber.Ctx) error {
	return c.Render("", "")
}

// GetById implements TreatmentHandler.
func (t *treatmentHandlerImpl) GetById(c *fiber.Ctx) error {
	return c.Render("", "")
}

func NewTreatmentHandlerImpl(ts treatment_service.TreatmentService, wg *sync.WaitGroup) TreatmentHandler {
	return &treatmentHandlerImpl{
		ts: ts,
		wg: wg,
	}
}
