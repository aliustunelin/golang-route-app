package main

import (
	"golang-route-app/app"
	"golang-route-app/configs"
	"golang-route-app/repos"
	"golang-route-app/routes"
	"golang-route-app/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// run setup.go start mongoDB
func main() {
	appRoute := fiber.New()

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

	routes.ApiRoute(appRoute, td)

	appRoute.Listen(":8080")

}
