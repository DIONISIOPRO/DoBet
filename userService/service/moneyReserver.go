package service

import (
	"sync"
)
type Reserve struct{
	Amount float64
	Hash string
}
type MoneyReserver struct {
	lock *sync.Mutex
	store map[string]Reserve
}

func (reserver MoneyReserver) ReserveMoney(userid string, money float64, hash string) {
	reserver.lock.Lock()
	defer reserver.lock.Unlock()
	account, ok := reserver.store[userid]
	if !ok {
		reserveMoney[userid] = Reserve{
			Amount: money,
			Hash: hash,
		}
	}
	account.Amount += money
	reserver.store[userid] = account
}

func (reserver *MoneyReserver) UnReserveMoney(userid string, money float64, hash string) {
	reserver.lock.Lock()
	defer reserver.lock.Unlock()
	account := Reserve{}
	for _, account := range reserver.store {
		if account.Hash == hash {
			account.Amount-= money
			break
		}
	}
	reserver.store[userid] = account
}
