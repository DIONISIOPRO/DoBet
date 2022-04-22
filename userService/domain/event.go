package domain

import "encoding/json"

const (
	USERLOGIN           = "use.login"
	USERLOGOUT          = "user.logout"
	USERDELETE          = "user.delete"
	USERCREATED         = "user.created"
	USERUPDATE          = "user.update"
	USERCONFIRMWITHDRAW = "user.confirm.withdraw"
	USERCONFIRMBET      = "user.confirm.bet"
	USERREQUESTWITHDRAW = "user.request.withdraw"
	USERREQUESTBET      = "user.request.bet"
	USERDEPOSIT         = "user.deposit"
	USERWITHDRAW        = "user.withdraw"
	USERBET             = "user.bet"
	USERWIN             = "user.win"
)

var QueuesToListenning = []string{
	USERDELETE, USERUPDATE,USERCREATED, USERREQUESTBET, USERREQUESTWITHDRAW, USERDEPOSIT, USERWIN, USERWITHDRAW, USERBET, USERLOGIN, USERLOGOUT}
var QueuesToPublish = []string{
	USERDELETE, USERUPDATE, USERCONFIRMWITHDRAW, USERCONFIRMBET, USERCREATED,
}

type (
	Event interface {
		ToByteArray() ([]byte, error)
	}

	UserCreatedEvent struct {
		User User `json:"user"`
	}

	UserDeletedEvent struct {
		UserId string `json:"user_id"`
	}
	UserUpdateEvent struct {
		UserId string `json:"user_id"`
		User   User   `json:"user"`
	}
	CheckMoneyEvent struct {
		UserId string  `json:"user_id"`
		Amount float64 `json:"amount"`
		Hash   string  `json:"hash"`
	}

	ConfirmMoneyEvent struct {
		Hash        string `json:"hash"`
		CanWithDraw bool   `json:"can_withdraw"`
	}

	AddMoneyEvent struct {
		UserId string  `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	SubtractMoneyEvent struct {
		UserId string  `json:"user_id"`
		Amount float64 `json:"amount"`
		Hash   string  `json:"hash"`
	}
)

func (event UserDeletedEvent) ToByteArray() ([]byte, error) {
	return json.Marshal(event)
}

func (event SubtractMoneyEvent) ToByteArray() ([]byte, error) {
	return json.Marshal(event)
}

func (event AddMoneyEvent) ToByteArray() ([]byte, error) {
	return json.Marshal(event)
}

func (event ConfirmMoneyEvent) ToByteArray() ([]byte, error) {
	return json.Marshal(event)
}

func (event CheckMoneyEvent) ToByteArray() ([]byte, error) {
	return json.Marshal(event)
}

func (event UserUpdateEvent) ToByteArray() ([]byte, error) {
	return json.Marshal(event)
}

func (event UserCreatedEvent) ToByteArray() ([]byte, error) {
	return json.Marshal(event)
}
