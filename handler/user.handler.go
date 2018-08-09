package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
}

func NewUserHandler(router *mux.Router) {

}

func (h *UserHandler) GetLoginPage(w http.ResponseWriter, r *http.Request) {

}
