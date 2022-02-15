package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Match struct {
	ID          primitive.ObjectID `bson:"_id"`
	Match_id    string             `json:"match_id" validate:"required"`
	Home_team   Team               `json:"home_team" validate:"required"`
	Away_team   Team               `json:"away_team" validate:"required"`
	Status      string             `json:"status" validate:"required"`
	Result      Match_Result       `json:"result" validate:"required"`
	Odds        Odds               `json:"odds" validate:"required"`
	IsProcessed bool               `json:"isprocessed" validate:"required"`
}
