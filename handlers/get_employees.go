package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-mongo-hrms/db"
	"go-fiber-mongo-hrms/types"
)

func GetEmployees(c *fiber.Ctx) error {
	employees := make([]types.Employee, 0)

	collection := db.Instance.Collection(db.EmployeesCollection)
	cur, err := collection.Find(c.Context(), bson.D{})
	if err != nil {
		log.Println("unable to fetch documents:", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	if err := cur.All(c.Context(), &employees); err != nil {
		log.Println("unable to fetch documents:", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(employees)
}
