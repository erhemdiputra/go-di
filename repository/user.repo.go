package repository

import (
	"github.com/erhemdiputra/exec-go-di/models"
)

type IUserRepo interface {
	GetList() []models.User
}

type UserRepo struct {
}

func NewUserRepo() IUserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetList() []models.User {
	list := []models.User{
		models.User{ID: 1, Name: "James"},
		models.User{ID: 2, Name: "John"},
		models.User{ID: 3, Name: "Robert"},
	}

	return list
}
