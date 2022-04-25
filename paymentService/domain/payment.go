package domain

var NotEnoughMoney = "not enougth money"

type Deposit struct {
	Phone_number string  `json:"phone_number" bson:"phone_number"`
	Amount float64 `json:"amount" bson:"amount"`
}
type WithDraw struct {
	Deposit
}
