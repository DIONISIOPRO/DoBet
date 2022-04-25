package domain

import "encoding/json"

const (
	USERDELETED        = "user.deleted"
	USERCREATED        = "user.created"
	USERUPDATED        = "user.update"
	USERBETCREATED     = "user.bet.created"
	USERBALANCEUPDATED = "user.balance.updated"
	USERBETPAYED       = "user.bet"
	USERBETWIN         = "user.bet.win"
)

var EventsToListenning = []string{
	USERDELETED, USERUPDATED, USERBETCREATED, USERBETWIN, USERCREATED,
}
var EventsToPublish = []string{
	USERBALANCEUPDATED, USERBETWIN,
}

type UserCreatedEvent struct {
	User User `json:"user"`
}

type BetCreatedEvent struct {
	User_id      string  `json:"user_id"`
	Bet_Id       string  `json:"bet_id"`
	Phone_number string  `json:"phone_number"`
	Amount       float64 `json:"amount"`
}

type UserWinEvent struct {
	Phone_number string  `json:"phone_number"`
	Amount       float64 `json:"amount"`
}

type BetPayedEvent struct {
	Bet_id string `json:"bet_id"`
}

type BetRefundEvent struct {
	Bet_Id       string  `json:"bet_id"`
	Phone_number string  `json:"phone_number"`
	Amount       float64 `json:"amount"`
}

type BalanceUpdateEvent struct {
	User_Id      string  `json:"bet_id"`
	Phone_number string  `json:"phone_number"`
	Amount       float64 `json:"amount"`
}

type UserDeletedEvent struct {
	UserId string `json:"user_id"`
}

type UserUpdateEvent struct {
	UserId string `json:"user_id"`
	User   User   `json:"user"`
}

func (ev *UserCreatedEvent) IsValid() bool {
	return ev.User.Account_balance == float64(0)
}

func (ev *UserUpdateEvent) IsValid() bool {
	return ev.User.User_id == ev.UserId
}

