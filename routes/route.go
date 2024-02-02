package routes

import (
	"golang-route-app/app"

	"github.com/gofiber/fiber/v2"
)

func ApiRoute(ctx *fiber.App, td app.LocationHandler) {

	ctx.Delete("/api/deleteAll", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.DeleteAllLocation(c)
	})

	ctx.Post("/api/createLocation", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.CreateLocation(c)
	})

	ctx.Get("/api/getLocations", func(c *fiber.Ctx) error {

		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.GetAllLocation(c)
	})

	//id parameters include golang IDs, not mongo ID
	ctx.Delete("/api/deleteLocation", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.DeleteLocation(c)
	})

	ctx.Get("/api/location", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.GetByNameWithDataLocation(c)
	})

	ctx.Post("/api/updateLocation", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.UpdateByIDLocation(c)
	})

	ctx.Post("/api/routing", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.RouteLocation(c)
	})

}
