package api

type pymentapi struct{

}

func NewPaymentApi() PaymentApi{
	return &pymentapi{}
}

func (p *pymentapi) Deposit(amount float64, userid string) error{
 return nil
}

func (p *pymentapi) Withdraw(amount float64, userid string) error{
return nil
}