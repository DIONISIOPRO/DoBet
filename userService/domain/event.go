package domain

type UserDeletedEvent struct {
	UserId string
}
type UserUpdateEvent struct {
	User User
}
type CheckMoneyEvent struct {
	UserId string
	Amount float64
	Hash   string
}

type ConfirmMoneyEvent struct {
	Hash        string
	CanWithDraw bool
}

type AddMoneyEvent struct {
	UserId string
	Amount float64
}

type SubtractMoneyEvent struct {
	UserId string
	Amount float64
	Hash   string
}