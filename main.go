package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"go-fiber-mongo-hrms/db"
	"go-fiber-mongo-hrms/handlers"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Post("/employees", handlers.NewEmployee)
	app.Get("/employees", handlers.GetEmployees)
	app.Get("/employees/:id", handlers.GetEmployee)
	app.Put("/employees/:id", handlers.UpdateEmployee)
	app.Delete("/employees/:id", handlers.DeleteEmployee)

	log.Fatalln(app.Listen(":3000"))
}
