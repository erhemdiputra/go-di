package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/erhemdiputra/go-di/service"
	"github.com/gorilla/securecookie"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	UserService service.IUserService
	MapTemplate map[string]*template.Template
}

func NewUserHandler(router *httprouter.Router, cookieHandler *securecookie.SecureCookie,
	mapTemplate map[string]*template.Template) {

	userService := service.NewUserService(cookieHandler)

	handler := UserHandler{
		UserService: userService,
		MapTemplate: mapTemplate,
	}

	router.GET("/login", handler.GetLoginPage)
	router.POST("/login", handler.PostLoginPage)
	router.POST("/logout", handler.Logout)
	router.GET("/home", handler.GetHomePage)
}

func (h *UserHandler) GetLoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := h.MapTemplate["login"].Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *UserHandler) PostLoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.FormValue("name")
	password := r.FormValue("password")

	if name == "" && password == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// TODO : Check DB//
	h.UserService.SetSession(name, w)
	http.Redirect(w, r, "/home", http.StatusFound)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.UserService.ClearSession(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (h *UserHandler) GetHomePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userName := h.UserService.GetUserName(r)

	if userName == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	data := map[string]interface{}{
		"username": userName,
	}

	h.MapTemplate["home"].Execute(w, data)
}
