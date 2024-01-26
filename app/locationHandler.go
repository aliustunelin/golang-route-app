package app

import (
	"golang-route-app/models"
	"golang-route-app/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationHandler struct {
	Service services.LocationService
}

func (h LocationHandler) CreateLocation(c *fiber.Ctx) error {
	var loction models.Location

	if err := c.BodyParser(&loction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.LocationInsert(loction)

	if err != nil || result.Status == false {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h LocationHandler) GetAllLocation(c *fiber.Ctx) error {
	result, err := h.Service.LocationGetAll()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h LocationHandler) DeleteLocation(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.LocationDelete(cnv)
	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}
