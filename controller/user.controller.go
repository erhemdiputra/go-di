package controller

import (
	"github.com/erhemdiputra/practice-mvc/models"
	"github.com/erhemdiputra/practice-mvc/service"
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
