package domain

import "encoding/json"

const (
	USERLOGIN   = "use.login"
	USERLOGOUT  = "user.logout"
	USERDELETE  = "user.delete"
	USERCREATED = "user.created"
	USERUPDATE  = "user.update"
)

var EventsToPublish = []string{USERLOGIN, USERLOGOUT}
var EventsToListenning = []string{USERUPDATE, USERDELETE, USERCREATED}

type (
	Event interface {
		ToByteArray() ([]byte , error)
	}
	DeleteUserEvent struct {
		UserId string `json:"user_id"`
	}
	AddUserEvent struct {
		User User `json:"user"`
	}
	UpdateUserEvent struct {
		UserId string `json:"user_id"`
		User   User   `json:"user"`
	}
	LoginEvent struct {
		UserId string `json:"user_id"`
	}
	LogOutEvent struct {
		UserId string `json:"user_id"`
	}
)

func (event LogOutEvent) ToByteArray() ([]byte , error){
	data, err := json.Marshal(event)
	return data, err

}

func (event LoginEvent) ToByteArray()([]byte , error) {
	data, err := json.Marshal(event)
	return data, err
}
