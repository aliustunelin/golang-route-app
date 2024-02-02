package services

import (
	"golang-route-app/dto"
	"golang-route-app/models"
	"golang-route-app/repos"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultLocationService struct {
	Repo repos.LocationRepos
}

type LocationService interface {
	InsertLocationService(Location models.Location) (*dto.LocationDTO, error)
	GetAllLocationService() ([]models.Location, error)
	DeleteLocationService(id primitive.ObjectID) (bool, error)
	GetByNameWithDataLocationService(id primitive.ObjectID) ([]models.Location, error)
	UpdateByIDLocationService(location models.Location) (*dto.LocationDTO, error)
	RouteLocationService(location models.Location) ([]primitive.M, error)
	DeleteAllLocationService() (bool, error)
}

func (t DefaultLocationService) InsertLocationService(Location models.Location) (*dto.LocationDTO, error) {
	var res dto.LocationDTO
	if len(Location.Name) < 2 {
		res.Status = false
		return &res, nil
	}
	result, err := t.Repo.Insert(Location)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dto.LocationDTO{Status: result}
	return &res, nil
}

func (t DefaultLocationService) GetAllLocationService() ([]models.Location, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultLocationService) DeleteLocationService(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func (t DefaultLocationService) DeleteAllLocationService() (bool, error) {
	result, err := t.Repo.DeleteAll()
	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func (t DefaultLocationService) GetByNameWithDataLocationService(id primitive.ObjectID) ([]models.Location, error) {
	result, err := t.Repo.GetByNameWithData(id)
	if err != nil {
		log.Fatalln(err)
	}
	return result, nil

}

func (t DefaultLocationService) UpdateByIDLocationService(location models.Location) (*dto.LocationDTO, error) {
	var res dto.LocationDTO
	if len(location.Name) < 2 {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.UpdateByID(location)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dto.LocationDTO{Status: result}
	return &res, nil
}

func (t DefaultLocationService) RouteLocationService(location models.Location) ([]primitive.M, error) {
	result, err := t.Repo.Routing(location)
	if err != nil {
		log.Fatalln(err)
	}
	return result, nil

}

func NewLocationService(Repo repos.LocationRepos) DefaultLocationService {
	return DefaultLocationService{Repo: Repo}
}
