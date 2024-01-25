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
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepoDb := repos.NewToDoReposDb(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepoDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	//golang'deki id mongo' id değil
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":8080")

}
