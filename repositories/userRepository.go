package repositories

import (
	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
)

type userRepository struct{
}

func NewUserRepository() UserRepository{
	return &userRepository{}
}

var userColletion = database.OpenCollection("users")

func (repo *userRepository) Deposit(amount float64, userid string) error {

	return nil

}

func (repo *userRepository) Withdraw(amount float64, userid string) error {
	return nil
}


func (repo *userRepository) Login(user models.User) error {
	return nil
}

func (repo *userRepository) SignUp(user models.User) error {

	return nil

}

func (repo *userRepository) Win(amount float64, user_id string) {
	
}

func (repo *userRepository) Users()([]models.User, error){
	return []models.User{}, nil
}

