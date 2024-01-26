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
// https://gist.github.com/aliustunelin/8e2d1fa132177507827f94d7f716795f
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
	appRoute.Post("/api/location/:id", td.GetByNameWithDataLocation)
	appRoute.Post("/api/locaation", td.UpdateByIDLocation)
	appRoute.Listen(":8080")

}
