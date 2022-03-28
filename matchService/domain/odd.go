package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Odds struct {
	ID                primitive.ObjectID `bson:"_id"`
	Odd_id            string             `bson:"odd_id" json:"odd_id"`
	Match_id          string             `bson:"match_id" json:"match_id"`
	WinnerMarketOdd   WinnerMarketOdd    `bson:"winner_odd" json:"winner_odd"`
	AllScoreMarketOdd AllScoreMarketOdd  `bson:"all_score_odd" json:"all_score_odd"`
}

type WinnerMarketOdd struct {
	Away float64 `bson:"away" json:"away"`
	Home float64 `bson:"home" json:"home"`
	Draw float64 `bson:"draw" json:"draw"`
}

type AllScoreMarketOdd struct {
	Yes float64 `bson:"yes" json:"yes"`
	No  float64 `bson:"no" json:"no"`
}
