package util

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		now := time.Now()
		defer func() {
			fmt.Printf(
				"method=%s, url=%s, host=%s, path=%s, duration=%s, status=%d\n",
				c.Method(),
				c.Request().URI().String(),
				c.Hostname(),
				c.Path(),
				time.Since(now).String(),
				c.Response().StatusCode(),
			)
		}()

		return c.Next()
	}
}
