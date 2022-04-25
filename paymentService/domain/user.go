package domain

type User struct {
	User_id         string  `json:"user_id"`
	Phone_number    string  `json:"phone_number"`
	Account_balance float64 `json:"account_balance"`
}
