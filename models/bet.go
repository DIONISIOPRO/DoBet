package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bet struct {
	ID            primitive.ObjectID `bson:"_id"`
	Bet_id        string             `json:"bet_id" validate:"required"`
	TotalAmount   float64            `json:"totalamount" validate:"required"`
	GlobalOdd     float64            `json:"totalodd" validate:"required"`
	Bet_owner     string             `json:"bet_owner" validate:"required"`
	Potencial_win float64            `json:"potencial_win" validate:"required"`
	IsFinished    bool               `json:"isfinished" validate:"required"`
	BetGroup      []SingleBet        `json:"betgroup" validate:"required"`
}

type SingleBet struct {
	Match_id    string      `json:"match_id" validate:"required"`
	IsProcessed bool        `json:"isprocessed" validate:"required"`
	Amount      float64     `json:"amount" validate:"required"`
	Market      interface{} `json:"market" validate:"required"`
	IsLose      bool        `json:"islose" validate:"required"`
	Option      BetOption   `json:"options" validate:"required"`
	Odd         float64     `json:"odd" validate:"required"`
}
