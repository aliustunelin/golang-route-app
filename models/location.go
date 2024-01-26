package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Lat         float64            `json:"lat,omitempty"`
	Lon         float64            `json:"lon,omitempty"`
	Name        string             `json:"name,omitempty"`
	MarkerColor string             `json:"markerColor,omitempty"`
}
