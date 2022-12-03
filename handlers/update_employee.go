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

func UpdateEmployee(c *fiber.Ctx) error {
	employee := &types.Employee{}
	id := c.Params("id")

	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("unable to parse bson id:", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	if err := c.BodyParser(employee); err != nil {
		log.Println("unable to parse body:", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	update := bson.M{}
	if len(employee.Name) > 0 {
		update["name"] = employee.Name
	}
	if employee.Age > 0 {
		update["age"] = employee.Age
	}
	if employee.Salary > 0 {
		update["salary"] = employee.Salary
	}

	if len(update) == 0 {
		log.Println("no field to update")
		return c.SendStatus(http.StatusBadRequest)
	}

	collection := db.Instance.Collection(db.EmployeesCollection)
	_, err = collection.UpdateOne(
		c.Context(),
		bson.M{"_id": bsonId},
		bson.M{
			"$set": update,
		},
	)
	if err != nil {
		log.Println("unable to update:", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	employee.ID = id
	return c.Status(http.StatusOK).JSON(employee)
}
