package cjson

import "github.com/gofiber/fiber/v2"

// JSONResponse adalah fungsi helper untuk membungkus respon JSON
// Mengatur status HTTP dan mengembalikan data JSON kepada client
func JSONResponse(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(data)
}
