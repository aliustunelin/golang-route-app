package main

import (
	"golang-route-app/app"
	"golang-route-app/configs"
	"golang-route-app/repos"
	"golang-route-app/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// run setup.go start mongoDB
func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	//mongodb name
	dbClient := configs.GetCollection(configs.DB, configs.EnvMongoName())

	TodoRepoDb := repos.NewLocationReposDb(dbClient)

	td := app.LocationHandler{Service: services.NewLocationService(TodoRepoDb)}

	limiterConfig := limiter.Config{
		Max:        20,
		Expiration: time.Second * 30,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
	}
	appRoute.Use(limiter.New(limiterConfig))

	appRoute.Post("/api/createLocation", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.CreateLocation(c)
	})

	appRoute.Get("/api/getLocations", func(c *fiber.Ctx) error {

		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.GetAllLocation(c)
	})

	//id parameters include golang IDs, not mongo ID
	appRoute.Delete("/api/deleteLocation", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.DeleteLocation(c)
	})

	appRoute.Get("/api/location", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.GetByNameWithDataLocation(c)
	})

	appRoute.Post("/api/updateLocation", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.UpdateByIDLocation(c)
	})

	appRoute.Post("/api/routing", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.RouteLocation(c)
	})
	appRoute.Listen(":8080")

}
