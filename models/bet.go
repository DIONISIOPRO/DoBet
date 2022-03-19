package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bet struct {
	ID            primitive.ObjectID `bson:"_id" json:"_"`
	Bet_id        string             `json:"bet_id" bson:"bet_id"`
	TotalAmount   float64            `json:"totalamount" validate:"required"`
	GlobalOdd     float64            `json:"globalodd"`
	Bet_owner     string             `json:"bet_owner" bson:"bet_owner" validate:"required"`
	Potencial_win float64            `json:"potencial_win"`
	IsFinished    bool               `json:"isfinished"`
	BetGroup      []SingleBet        `json:"betgroup" validate:"required"`
}

type SingleBet struct {
	Match_id    string      `json:"match_id" validate:"required"`
	IsProcessed bool        `json:"isprocessed"`
	Market      interface{} `json:"market" validate:"required"`
	IsLose      bool        `json:"islose"`
	Option      BetOption   `json:"options" validate:"required"`
	Odd         float64     `json:"odd" validate:"required"`
}
