package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type League struct {
	ID          primitive.ObjectID `bson:"_id"`
	League_id   string             `json:"league_id" validate:"required"`
	Name        string             `json:"name" validate:"required"`
	CountryName string             `json:"country"`
	Logo_url    string             `json:"logo_urls"`
}