package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

func CreateBet(owner models.User, market models.Market, option models.BetOption, amount float64) error {
	return repositories.CreateBet(owner, market, option,amount)

}