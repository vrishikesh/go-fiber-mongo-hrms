package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-fiber-mongo-hrms/db"
	"go-fiber-mongo-hrms/types"
)

func NewEmployee(c *fiber.Ctx) error {
	employee := &types.Employee{}

	if err := c.BodyParser(employee); err != nil {
		log.Println("unable to parse body:", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	if len(employee.Name) == 0 {
		log.Println("employee name is required")
		return c.Status(http.StatusBadRequest).SendString("employee name is required")
	}

	employee.ID = ""

	collection := db.Instance.Collection(db.EmployeesCollection)
	result, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		log.Println("unable to insert document:", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	bsonId := result.InsertedID.(primitive.ObjectID)
	employee.ID = bsonId.Hex()
	return c.Status(http.StatusCreated).JSON(employee)
}
