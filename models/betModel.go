package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bet struct {
	ID            primitive.ObjectID `bson:"_id"`
	Bet_id        string             `json:"bet_id" validate:"required"`
	Bet_owner     string             `json:"bet_owner" validate:"required"`
	Match_id      []string           `json:"match_id" validate:"required"`
	Amount        float64            `json:"amount" validate:"required"`
	Markets       []Market           `json:"market_id" validate:"required"`
	Options       []BetOption        `json:"options" validate:"required"`
	Odd           float64            `json:"odd" validate:"required"`
	Potencial_win float64            `json:"potencial_win" validate:"required"`
	IsProcessed   bool               `json:"isprocessed" validate:"required"`
	IsFinished    bool               `json:"isfinished" validate:"required"`
	RemainMatches int                `json:"remain_matches" validate:"required,min=1"`
	IsLose        bool               `json:"islose" validate:"required"`
}
