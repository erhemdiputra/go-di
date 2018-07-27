package handler

import (
	"net/http"

	"github.com/erhemdiputra/go-di/controller"
	"github.com/erhemdiputra/go-di/repository"
	"github.com/erhemdiputra/go-di/service"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserController *controller.UserController
	Router         *mux.Router
}

func NewUserHandler(router *mux.Router) *UserHandler {
	userRepo := repository.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &UserHandler{
		UserController: userController,
		Router:         router,
	}
}

func (h *UserHandler) Serve() {
	h.Router.Handle("/api/user/list", HandlerFunc(h.GetList)).Methods("GET")
}

func (h *UserHandler) GetList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	users := h.UserController.GetList()
	return users, nil
}
