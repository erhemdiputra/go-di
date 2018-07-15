package handler

import (
	"encoding/json"
	"net/http"

	"github.com/erhemdiputra/go-di/controller"
	"github.com/erhemdiputra/go-di/repository"
	"github.com/erhemdiputra/go-di/service"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	UserController *controller.UserController
	Router         *httprouter.Router
}

func NewUserHandler(router *httprouter.Router) *UserHandler {
	userRepo := repository.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &UserHandler{
		UserController: userController,
		Router:         router,
	}
}

func (h *UserHandler) Serve() {
	h.Router.GET("/user/list", h.GetList)
}

func (h *UserHandler) GetList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users := h.UserController.GetList()

	encoded, _ := json.Marshal(users)
	w.Write(encoded)
}
