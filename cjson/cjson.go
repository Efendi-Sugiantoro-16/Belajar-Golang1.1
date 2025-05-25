package cjson

import "github.com/gofiber/fiber/v2"

// JSONResponse is a helper function to return JSON response with status code and data
func JSONResponse(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(data)
}
