package services

import "github.com/alina-boichenko/data-users-api_gin-gonic.git/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
