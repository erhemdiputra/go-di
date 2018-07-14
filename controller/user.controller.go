package controller

import (
	"github.com/erhemdiputra/exec-go-di/models"
	"github.com/erhemdiputra/exec-go-di/service"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) GetList() []models.User {
	return c.UserService.GetList()
}
