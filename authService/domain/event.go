package domain

type (
	UserDeletedEvent struct {
		UserId string
	}
	UserUpdateEvent struct {
		User User
	}
	CheckMoneyEvent struct {
		UserId string
		Amount float64
		Hash   string
	}

	ConfirmMoneyEvent struct {
		Hash        string
		CanWithDraw bool
	}

	AddMoneyEvent struct {
		UserId string
		Amount float64
	}

	SubtractMoneyEvent struct {
		UserId string
		Amount float64
		Hash   string
	}
)
