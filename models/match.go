package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Match struct {
	ID           primitive.ObjectID `bson:"_id"`
	Match_id     string             `json:"match_id" validate:"required"`
	Home_team_id string             `json:"home_team_id" validate:"required"`
	Away_team_id string             `json:"away_team_id" validate:"required"`
	LeagueId     string             `json:"league" validate:"required"`
	Status       string             `json:"status" validate:"required"`
	Result       Match_Result       `json:"result" validate:"required"`
	Time         int64              `json:"time" validate:"required"`
	Odds         Odds               `json:"odds" validate:"required"`
}
