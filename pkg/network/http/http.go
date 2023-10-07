package http

import (
	"encoding/json"

	"github.com/LCRERGO/drucssGO/pkg/domain/entity"
	"github.com/LCRERGO/drucssGO/pkg/domain/service"
	"github.com/gofiber/fiber/v2"
)

type RouterFunc func(c *fiber.Ctx) error

type Handler interface {
	CreateUser() RouterFunc
	ReadAllUsers() RouterFunc
	ReadUser() RouterFunc
	UpdateUser() RouterFunc
	DeleteUser() RouterFunc
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateUser() RouterFunc {
	return func(c *fiber.Ctx) error {
		var inUser entity.User

		c.Accepts("application/json")
		if err := c.BodyParser(&inUser); err != nil {
			return err
		}

		if err := h.service.CreateUser(inUser); err != nil {
			return err
		}

		return nil
	}
}

func (h *handler) ReadAllUsers() RouterFunc {
	return func(c *fiber.Ctx) error {
		users, err := h.service.ReadAllUsers()
		if err != nil {
			return err
		}

		usersMarshalled, err := json.Marshal(users)
		if err != nil {
			return err
		}

		return c.Send(usersMarshalled)
	}
}

func (h *handler) ReadUser() RouterFunc {
	return func(c *fiber.Ctx) error {
		var inUser entity.User

		c.Accepts("application/json")
		if err := c.BodyParser(&inUser); err != nil {
			return err
		}

		user, err := h.service.ReadUser(inUser)
		if err != nil {
			return err
		}

		userMarshalled, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return c.Send(userMarshalled)
	}
}

func (h *handler) UpdateUser() RouterFunc {
	return func(c *fiber.Ctx) error {
		var inUser entity.User

		c.Accepts("application/json")
		if err := c.BodyParser(&inUser); err != nil {
			return err
		}
		err := h.service.UpdateUser(inUser)
		if err != nil {
			return err
		}

		return nil
	}
}

func (h *handler) DeleteUser() RouterFunc {
	return func(c *fiber.Ctx) error {
		var inUser entity.User

		c.Accepts("application/json")
		if err := c.BodyParser(&inUser); err != nil {
			return err
		}
		err := h.service.DeleteUser(inUser)
		if err != nil {
			return err
		}

		return nil
	}
}

func Configure(handler Handler, fiberConn *fiber.App) {
	fiberConn.Post("/v1/user", handler.CreateUser())
	fiberConn.Get("/v1/users", handler.ReadAllUsers())
	fiberConn.Get("/v1/user", handler.ReadUser())
	fiberConn.Patch("/v1/user", handler.UpdateUser())
	fiberConn.Delete("/v1/user", handler.DeleteUser())
}
