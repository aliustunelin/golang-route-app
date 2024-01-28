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

// projeyi mainden ayağa kaldıraz
// sertup.go'daki connectedDB'yi çalıştıracağız
// https://gist.github.com/aliustunelin/8e2d1fa132177507827f94d7f716795f
func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	//mongodb name
	dbClient := configs.GetCollection(configs.DB, "YukatLocations")

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

	//golang'deki id mongo' id değil
	appRoute.Delete("/api/deleteLocation/:id", func(c *fiber.Ctx) error {
		if c.Locals("limit") != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}
		return td.DeleteLocation(c)
	})

	appRoute.Get("/api/location/:id", func(c *fiber.Ctx) error {
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
		return td.RoutingLocation(c)
	})
	appRoute.Listen(":8080")

}
