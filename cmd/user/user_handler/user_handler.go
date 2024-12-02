package user_handler

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
}
