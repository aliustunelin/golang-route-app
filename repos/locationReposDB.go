package repos

import (
	"context"
	"errors"
	"golang-route-app/models"
	"log"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks//repository/mockLocationRepos.go -package=repos golang-route-app/repos LocationRepos
type LocationReposDB struct {
	LocationCollection *mongo.Collection
}

type LocationRepos interface {
	Insert(location models.Location) (bool, error)
	GetAll() ([]models.Location, error)
	Delete(id primitive.ObjectID) (bool, error)
	GetByNameWithData(id primitive.ObjectID) ([]models.Location, error)
	UpdateByID(location models.Location) (bool, error)
	Routing(location models.Location) ([]primitive.M, error)
	DeleteAll() (bool, error)
}

func (t LocationReposDB) Insert(location models.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	location.Id = primitive.NewObjectID()

	result, err := t.LocationCollection.InsertOne(ctx, location)

	if result.InsertedID == nil || err != nil {
		errors.New("failed add!!")
		return false, err
	}

	return true, nil
}

func (t LocationReposDB) DeleteAll() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.LocationCollection.DeleteMany(ctx, bson.M{})
	if err != nil || result.DeletedCount == 0 {
		log.Fatalln(err)
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

func (t LocationReposDB) UpdateByID(location models.Location) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        location.Name,
			"lat":         location.Lat,
			"lon":         location.Lat,
			"markerColor": location.MarkerColor,
		},
	}

	filter := bson.M{
		"id": location.Id,
	}

	result := t.LocationCollection.FindOneAndUpdate(ctx, filter, update)
	if result == nil {
		log.Fatalln(result, " is not update!")
	}

	return true, nil
}

func (t LocationReposDB) Routing(reqLocation models.Location) ([]primitive.M, error) {
	var ilocation models.Location
	var ilocations []primitive.M

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.LocationCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}

	for result.Next(ctx) {
		if err := result.Decode(&ilocation); err != nil {
			log.Fatalln(err)
		}

		//route json create
		air_distance := math.Sqrt((ilocation.Lat-reqLocation.Lat)*(ilocation.Lat-reqLocation.Lat) + (ilocation.Lon-reqLocation.Lon)*(ilocation.Lon-reqLocation.Lon))
		testFilter := bson.M{
			"id":                  ilocation.Id,
			"target_location_lon": reqLocation.Lon,
			"target_location_lat": reqLocation.Lat,
			"lon":                 ilocation.Lon,
			"lat":                 ilocation.Lat,
			"name":                ilocation.Name,
			"air_distance":        air_distance,
		}
		ilocations = append(ilocations, testFilter)
	}

	//sorting
	temp := ilocations[0]["air_distance"]
	for j, select_location_first := range ilocations {
		temp = ilocations[j]["air_distance"]

		log.Println(select_location_first)
		for i, select_location := range ilocations {

			if temp.(float64) < select_location["air_distance"].(float64) {
				tempLocation := ilocations[i]
				ilocations[i] = ilocations[j]
				ilocations[j] = tempLocation
			}
		}

	}

	return ilocations, nil

}
func NewLocationReposDb(dbClient *mongo.Collection) LocationReposDB {
	return LocationReposDB{LocationCollection: dbClient}
}
