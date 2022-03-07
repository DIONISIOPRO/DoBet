package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID          primitive.ObjectID `bson:"_id"`
	Team_id     string             `json:"team_id" validate:"required"`
	Name        string             `json:"name" validate:"required"`
	Logo_url    string             `json:"logo_url"`
	CountryName string             `json:"country" validate:"required"`
}
