package api

type MpesaApi struct {
	Key string
}

func (m *MpesaApi) Deposit(amount float64, userid string) error {
	return nil
}

func (m *MpesaApi) Withdraw(amount float64, userid string) error {
	return nil
}
