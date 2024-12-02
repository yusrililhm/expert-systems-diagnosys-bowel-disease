package user_handler

import (
	"sync"

	"healthy-bowel/cmd/user/user_service"

	"github.com/gofiber/fiber/v2"
)

type userHandlerImpl struct {
	us user_service.UserService
	wg *sync.WaitGroup
}

// ChangePassword implements UserHandler.
func (uh *userHandlerImpl) ChangePassword(c *fiber.Ctx) error {
	return c.Render("", "")
}

// Edit implements UserHandler.
func (uh *userHandlerImpl) Edit(c *fiber.Ctx) error {
	return c.Render("", "")
}

// Login implements UserHandler.
func (uh *userHandlerImpl) Login(c *fiber.Ctx) error {
	return c.Render("", "")
}

// Profile implements UserHandler.
func (uh *userHandlerImpl) Profile(c *fiber.Ctx) error {
	return c.Render("", "")
}

// Register implements UserHandler.
func (uh *userHandlerImpl) Register(c *fiber.Ctx) error {
	return c.Render("", "")
}

func NewUserHandlerImpl(us user_service.UserService, wg *sync.WaitGroup) UserHandler {
	return &userHandlerImpl{
		us: us,
		wg: wg,
	}
}
