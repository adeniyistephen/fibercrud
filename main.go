package main

import (
	"fiberCRUD/database"
	"fiberCRUD/route"

	"github.com/gofiber/fiber/v2"
)


/*
	TODO: make the api a little bit more production friendly and you can do more on the testing as well, well done.
*/


func main() {
	app := fiber.New()

	database.ConnectDb()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo-List-API :)")
	})

	app.Get("/tasks", route.GetAllTasks)

	app.Get("/task/:id", route.GetTask)

	app.Post("/add_task", route.AddTask) //Post endpoint for add a new task.

	app.Delete("/delete_task/:id", route.DeleteTask) //Delete endpoint for removing an existing task.

	app.Patch("/update_task/:id", route.UpdateTask) //Patch endpoint for updating an existing task.


	app.Listen(":8000")
}