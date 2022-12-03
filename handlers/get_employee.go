package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-fiber-mongo-hrms/db"
	"go-fiber-mongo-hrms/types"
)

func GetEmployee(c *fiber.Ctx) error {
	employee := &types.Employee{}
	id := c.Params("id")

	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("unable to parse bson:", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	collection := db.Instance.Collection(db.EmployeesCollection)
	result := collection.FindOne(c.Context(), bson.M{"_id": bsonId})
	if err := result.Err(); err != nil {
		log.Println("unable to find document:", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	err = result.Decode(employee)
	if err != nil {
		log.Println("unable to decode document:", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(employee)
}
