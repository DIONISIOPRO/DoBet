package service

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var OddServicee = &oddService{}
type oddService struct {
	repository repository.OddRepository
}

func (service *oddService) UpSertOdd(odd models.Odds) error{
	return service.repository.UpSertOdd(odd)

}
func (service *oddService) DeleteOdd(odd_id string) error{
return service.repository.DeleteOdd(odd_id)
}
