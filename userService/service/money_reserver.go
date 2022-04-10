package service

import (
	"sync"
)

type Reserve struct {
	Amount float64
	Hash   string
}
type MoneyReserver struct {
	lock  *sync.Mutex
	store map[string]Reserve
}

func newMoneyReserver(lock *sync.Mutex) *MoneyReserver {
	store := make(map[string]Reserve)
	return &MoneyReserver{
		lock:  lock,
		store: store,
	}

}

func (reserver *MoneyReserver) ReserveMoney(userid string, money float64, hash string) {
	reserver.lock.Lock()
	defer reserver.lock.Unlock()
	account, ok := reserver.store[userid]
	if !ok {
		reserver.store[userid] = Reserve{
			Amount: money,
			Hash:   hash,
		}
	}
	account.Amount += money
	reserver.store[userid] = account
}

func (reserver *MoneyReserver) UnReserveMoney(userid string, money float64, hash string) {
	reserver.lock.Lock()
	defer reserver.lock.Unlock()
	account := reserver.store[userid]
	account.Amount = account.Amount - money
	reserver.store[userid] = account
}

func (reserver *MoneyReserver) GetReservedMoneyByUserId(userId string) float64 {
	account, ok := reserver.store[userId]
	if !ok {
		return -1
	}
	return account.Amount
}
