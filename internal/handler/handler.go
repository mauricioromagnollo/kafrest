package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Handle(c *fiber.Ctx) error
}
