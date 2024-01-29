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

	result, err := h.Service.InsertLocationService(loction)

	if err != nil || result.Status == false {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h LocationHandler) GetAllLocation(c *fiber.Ctx) error {
	result, err := h.Service.GetAllLocationService()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h LocationHandler) DeleteLocation(c *fiber.Ctx) error {
	var loction models.Location

	if err := c.BodyParser(&loction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	cnv, _ := primitive.ObjectIDFromHex(loction.Id.Hex())

	// get ID from url or bodyParser, first I prefer url then move the jsonParser
	result, err := h.Service.DeleteLocationService(cnv)
	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}

func (h LocationHandler) GetByNameWithDataLocation(c *fiber.Ctx) error {
	var loction models.Location

	if err := c.BodyParser(&loction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	cnv, _ := primitive.ObjectIDFromHex(loction.Id.Hex())

	result, err := h.Service.GetByNameWithDataLocationService(cnv)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h LocationHandler) UpdateByIDLocation(c *fiber.Ctx) error {

	var loction models.Location

	if err := c.BodyParser(&loction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.UpdateByIDLocationService(loction)

	if err != nil || result.Status == false {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h LocationHandler) RouteLocation(c *fiber.Ctx) error {
	var location models.Location

	if err := c.BodyParser(&location); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	result, err := h.Service.RouteLocationService(location)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)

}
