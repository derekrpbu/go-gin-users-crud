package services

import "derek-api/models"

// interface to define service/api contracts
// we will implement these function using a Struct

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}


