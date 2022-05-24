package domain

var NotEnoughMoney = "not enougth money"

type Deposit struct {
	User_Id string  `json:"user_id" bson:"user_id"`
	Amount float64 `json:"amount" bson:"amount"`
}
type WithDraw struct {
	User_id string  `json:"user_id" bson:"user_id"`
	Amount float64 `json:"amount" bson:"amount"`
}
