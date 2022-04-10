package service

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReserveMoney(t *testing.T) {
	moneyReserver := newMoneyReserver(&sync.Mutex{})
	var teststore = make(map[string]Reserve)
	moneyReserver.store = teststore
	moneyReserver.ReserveMoney("id", float64(10), "hash")
	money, ok := teststore["id"]
	assert.Equal(t, float64(10), money.Amount)
	assert.Equal(t, true, ok)
}

func TestUnreserveMoney(t *testing.T) {
	moneyReserver := newMoneyReserver(&sync.Mutex{})
	var teststore = make(map[string]Reserve)
	moneyReserver.store = teststore
	moneyReserver.ReserveMoney("id", float64(10), "hash")
	money, ok := teststore["id"]
	assert.Equal(t, float64(10), money.Amount)
	assert.Equal(t, true, ok)
	moneyReserver.UnReserveMoney("id", float64(10), "hash")
	money, ok = teststore["id"]
	assert.Equal(t, float64(0), money.Amount)
	assert.Equal(t, true, ok)
}

func TestGetByID(t *testing.T) {
	moneyReserver := newMoneyReserver(&sync.Mutex{})
	var teststore = make(map[string]Reserve)
	moneyReserver.store = teststore
	moneyReserver.ReserveMoney("id", float64(10), "hash")
	money, ok := teststore["id"]
	assert.Equal(t, float64(10), money.Amount)
	assert.Equal(t, true, ok)
	money1 := moneyReserver.GetReservedMoneyByUserId("id")
	assert.Equal(t, float64(10), money1)
	moneyReserver.UnReserveMoney("id", float64(5), "hash")
	money, ok = teststore["id"]
	assert.Equal(t, float64(5), money.Amount)
	assert.Equal(t, true, ok)
	money2 := moneyReserver.GetReservedMoneyByUserId("id")
	assert.Equal(t, float64(5), money2)
	money2 = moneyReserver.GetReservedMoneyByUserId("iid")
	assert.Equal(t, float64(-1), money2)
}
