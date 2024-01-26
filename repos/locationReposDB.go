package repos

import (
	"context"
	"errors"
	"golang-route-app/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationReposDB struct {
	LocationCollection *mongo.Collection
}

type LocationRepos interface {
	Insert(location models.Location) (bool, error)
	GetAll() ([]models.Location, error)
	Delete(id primitive.ObjectID) (bool, error)
	GetByNameWithData(id primitive.ObjectID) ([]models.Location, error)
}

func (t LocationReposDB) Insert(location models.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	location.Id = primitive.NewObjectID()

	result, err := t.LocationCollection.InsertOne(ctx, location)

	if result.InsertedID == nil || err != nil {
		errors.New("failed add")
		return false, err
	}

	return true, nil
}

func (t LocationReposDB) GetAll() ([]models.Location, error) {
	var ilocation models.Location
	var ilocations []models.Location

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.LocationCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&ilocation); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		ilocations = append(ilocations, ilocation)
	}

	return ilocations, nil

}

func (t LocationReposDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.LocationCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}

	return true, nil
}

func (t LocationReposDB) GetByNameWithData(id primitive.ObjectID) ([]models.Location, error) {
	var ilocation models.Location
	var ilocations []models.Location

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.LocationCollection.Find(ctx, bson.M{"id": id})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&ilocation); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		ilocations = append(ilocations, ilocation)
	}

	return ilocations, nil
}

func NewLocationReposDb(dbClient *mongo.Collection) LocationReposDB {
	return LocationReposDB{LocationCollection: dbClient}
}
