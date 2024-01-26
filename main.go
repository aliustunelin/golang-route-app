package main

import (
	"golang-route-app/app"
	"golang-route-app/configs"
	"golang-route-app/repos"
	"golang-route-app/services"

	"github.com/gofiber/fiber/v2"
)

// projeyi mainden ayağa kaldıraz
// sertup.go'daki connectedDB'yi çalıştıracağız
func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	//mongodb name
	dbClient := configs.GetCollection(configs.DB, "YukatLocations")

	TodoRepoDb := repos.NewLocationReposDb(dbClient)

	td := app.LocationHandler{Service: services.NewLocationService(TodoRepoDb)}

	appRoute.Post("/api/location", td.CreateLocation)
	appRoute.Get("/api/locations", td.GetAllLocation)
	//golang'deki id mongo' id değil
	appRoute.Delete("/api/location/:id", td.DeleteLocation)
	appRoute.Listen(":8080")

}
