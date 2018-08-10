package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
}

func NewUserHandler(router *httprouter.Router) {

}

func (h *UserHandler) GetLoginPage(w http.ResponseWriter, r *http.Request) {

}
