package server

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
	if err != nil {
		return err
	}
	return nil
}
