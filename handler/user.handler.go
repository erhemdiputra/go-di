package handler

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/erhemdiputra/go-di/repository"
	"github.com/erhemdiputra/go-di/service"
	"github.com/gorilla/securecookie"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	UserService service.IUserService
	MapTemplate map[string]*template.Template
}

func NewUserHandler(router *httprouter.Router, db *sql.DB,
	cookieHandler *securecookie.SecureCookie, mapTemplate map[string]*template.Template) {

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(cookieHandler, userRepo)

	handler := UserHandler{
		UserService: userService,
		MapTemplate: mapTemplate,
	}

	router.GET("/login", handler.GetLoginPage)
	router.POST("/login", wrapHandler(handler.PostLoginPage))
	router.POST("/logout", handler.Logout)
	router.GET("/home", handler.GetHomePage)
}

func (h *UserHandler) GetLoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userName := h.UserService.GetUserName(r)
	if userName != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	err := h.MapTemplate["login"].Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *UserHandler) PostLoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" && password == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	ctx := r.Context()
	user, err := h.UserService.IsValidUser(ctx, username, password)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	h.UserService.SetSession(user.FullName, w)
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
