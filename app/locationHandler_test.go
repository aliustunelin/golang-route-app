package app

import (
	services "golang-route-app/mocks/service"
	"golang-route-app/models"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var td LocationHandler
var mockService *services.MockLocationService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = services.NewMockLocationService(ctrl)

	td = LocationHandler{mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_GetAllLocation(t *testing.T) {
	t.Skip("This test is not currently working!")
	teardown := setup(t)
	defer teardown()

	router := fiber.New()

	router.Get("/api/todos", td.GetAllLocation)

	var FakeDataForHandler = []models.Location{
		{Id: primitive.NewObjectID(), Name: "İstanbul Üniversitesi", MarkerColor: "blue", Lon: 28.6543812, Lat: 66.7349137},
		{Id: primitive.NewObjectID(), Name: "Ankara Üniversitesi", MarkerColor: "yellow", Lon: 92.1245311, Lat: 12.6341839},
		{Id: primitive.NewObjectID(), Name: "İzmir Üniversitesi", MarkerColor: "green", Lon: 52.8425945, Lat: 72.2413353},
		{Id: primitive.NewObjectID(), Name: "Antalya Üniversitesi", MarkerColor: "brown", Lon: 48.2326346, Lat: 173.3123749},
	}

	mockService.EXPECT().LocationGetAll().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	resp, _ := router.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}
