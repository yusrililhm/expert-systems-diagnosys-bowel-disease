package symptom_handler

import (
	"github.com/gofiber/fiber/v2"
)

type SymptomHandler interface {
	Add(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}
