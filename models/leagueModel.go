package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type League struct {
	Response string `json:"response"`
	ID        primitive.ObjectID `bson:"_id"`
	League_id string             `json:"league_id" validate:"required"`
	Name      string             `json:"name" validate:"required"`
	Country   Country             `json:"country"`
	Logo_url  string             `json:"logo_urls"`
}

type Country struct{
	Name string
	Code string
	Flag string
}