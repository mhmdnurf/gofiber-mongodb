package controllers

import (
	"context"
	"gofiber-mongodb/configs"
	"gofiber-mongodb/models"
	"gofiber-mongodb/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bahanCollection *mongo.Collection = configs.GetCollection(configs.DB, "master-bahan")

func GetAllBahan(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := bahanCollection.Find(ctx, primitive.D{{}})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BahanResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	var bahan []models.Bahan = make([]models.Bahan, 0)
	for cursor.Next(ctx) {
		var b models.Bahan
		cursor.Decode(&b)
		bahan = append(bahan, b)
	}

	return c.Status(http.StatusOK).JSON(responses.BahanResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": bahan}})
}
