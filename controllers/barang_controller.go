package controllers

import (
	"context"
	"gofiber-mongodb/configs"
	"gofiber-mongodb/models"
	"gofiber-mongodb/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var barangCollection *mongo.Collection = configs.GetCollection(configs.DB, "master-barang")

var validate = validator.New()

func CreateBarang(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var barang models.Barang
	defer cancel()

	if err := c.BodyParser(&barang); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BarangResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&barang); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BarangResponse{Status: http.StatusBadRequest, Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newBarang := models.Barang{
		Id:           primitive.NewObjectID(),
		NamaBarang:   barang.NamaBarang,
		JumlahBarang: barang.JumlahBarang,
		Satuan:       barang.Satuan,
		Keterangan:   barang.Keterangan,
	}

	result, err := barangCollection.InsertOne(ctx, newBarang)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BarangResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.BarangResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result.InsertedID}})
}

func GetAllBarang(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := barangCollection.Find(ctx, primitive.D{{}})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BarangResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	var barang []models.Barang = make([]models.Barang, 0)
	for cursor.Next(ctx) {
		var b models.Barang
		cursor.Decode(&b)
		barang = append(barang, b)
	}

	return c.Status(http.StatusOK).JSON(responses.BarangResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": barang}})
}
