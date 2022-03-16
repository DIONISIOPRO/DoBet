package models

type Deposit struct {
	UserId string `json:"id" bson:"user_id"`
	Amount float64 `json:"amount" bson:"amount"`
}

type WithDraw struct {
	UserId string `json:"id" bson:"user_id"`
	Amount float64 `json:"amount" bson:"amount"`
}