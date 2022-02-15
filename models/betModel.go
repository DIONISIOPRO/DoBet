package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bet struct {
	ID            primitive.ObjectID `bson:"_id"`
	Bet_id        string             `json:"bet_id" validate:"required"`
	Bet_owner     User               `json:"bet_owner" validate:"required"`
	Match_id      []string           `json:"match_id" validate:"required"`
	Amount        float64            `json:"amount" validate:"required"`
	Market        Market             `json:"market_id" validate:"required"`
	Option        BetOption          `json:"option" validate:"required"`
	Odd           float64            `json:"odd" validate:"required"`
	Potencial_win float64            `json:"potencial_win" validate:"required"`
	IsProcessed   bool               `json:"isprocessed" validate:"required"`
	IsFinished    bool               `json:"isfinished" validate:"required"`
	Status        string             `json:"status" validate:"required"`
}
