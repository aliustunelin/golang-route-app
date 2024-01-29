package services

import (
	repos "golang-route-app/mocks/repository"
	"golang-route-app/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockRepo *repos.MockLocationRepos
var service LocationService

var fakeData = []models.Location{
	{Id: primitive.NewObjectID(), Name: "İstanbul Üniversitesi", MarkerColor: "blue", Lon: 28.6543812, Lat: 66.7349137},
	{Id: primitive.NewObjectID(), Name: "Ankara Üniversitesi", MarkerColor: "yellow", Lon: 92.1245311, Lat: 12.6341839},
	{Id: primitive.NewObjectID(), Name: "İzmir Üniversitesi", MarkerColor: "green", Lon: 52.8425945, Lat: 72.2413353},
	{Id: primitive.NewObjectID(), Name: "Antalya Üniversitesi", MarkerColor: "brown", Lon: 48.2326346, Lat: 173.3123749},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repos.NewMockLocationRepos(ct)
	service = NewLocationService(mockRepo)
	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultLocationService_GetAllLocationService(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(fakeData, nil)

	result, err := service.GetAllLocationService()
	if err != nil {
		t.Error(err)
	}
	//test funcs
	assert.NotEmpty(t, result)
}

func TestDefaultLocationService_InsertLocationService(t *testing.T) {
	td := setup(t)
	defer td()

	locationMini := models.Location{
		Lat:         35.1230753,
		Lon:         39.9578864,
		Name:        "Stanford",
		MarkerColor: "purple",
	}

	mockRepo.EXPECT().Insert(locationMini).Return(true, nil)

	result, err := service.InsertLocationService(locationMini)
	if err != nil {
		t.Error(err)
	}
	//test funcs
	assert.NotEmpty(t, result)
}

func TestDefaultLocationService_DeleteLocationService(t *testing.T) {
	td := setup(t)
	defer td()

	fakeID := fakeData[0].Id

	mockRepo.EXPECT().GetAll().Return(fakeData, nil)
	mockRepo.EXPECT().Delete(fakeID).Return(true, nil)

	result, err := service.DeleteLocationService(fakeID)
	if err != nil {
		t.Error(err)
	}
	//test funcs
	assert.Equal(t, result, true)
}

func TestDefaultLocationService_GetByNameWithDataLocationService(t *testing.T) {
	td := setup(t)
	defer td()

	fakeID := fakeData[0].Id

	result := mockRepo.EXPECT().GetByNameWithData(fakeID)

	assert.NotEmpty(t, result)
}
