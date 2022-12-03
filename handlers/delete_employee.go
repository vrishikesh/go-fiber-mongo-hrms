package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-fiber-mongo-hrms/db"
)

func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("unable to parse bson id:", err)
		return c.SendStatus(http.StatusBadGateway)
	}

	collection := db.Instance.Collection(db.EmployeesCollection)
	_, err = collection.DeleteOne(c.Context(), bson.M{"_id": bsonId})
	if err != nil {
		log.Println("unable to delete document:", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusNoContent)
}