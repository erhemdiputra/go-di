package handler

import (
	"errors"
	"net/http"

	"github.com/erhemdiputra/go-di/controller"
	"github.com/erhemdiputra/go-di/repository"
	"github.com/erhemdiputra/go-di/service"
)

type UserHandler struct {
	UserController *controller.UserController
}

func NewUserHandler() *UserHandler {
	userRepo := repository.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &UserHandler{
		UserController: userController,
	}
}

func (h *UserHandler) Serve() {
	http.Handle("/api/user/list", HandlerFunc(h.GetList))
}

func (h *UserHandler) GetList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != http.MethodGet {
		return nil, errors.New("Invalid Request")
	}

	users := h.UserController.GetList()
	return users, nil
}
