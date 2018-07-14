package service

import (
	"github.com/erhemdiputra/exec-go-di/models"
	"github.com/erhemdiputra/exec-go-di/repository"
)

type IUserService interface {
	GetList() []models.User
}

type UserService struct {
	UserRepo repository.IUserRepo
}

func NewUserService(userRepo repository.IUserRepo) IUserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) GetList() []models.User {
	return s.UserRepo.GetList()
}
