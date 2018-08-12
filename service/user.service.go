package service

import (
	"context"
	"net/http"

	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/repository"
	"github.com/gorilla/securecookie"
)

type IUserService interface {
	GetUserName(request *http.Request) string
	SetSession(userName string, response http.ResponseWriter) error
	ClearSession(response http.ResponseWriter)
	IsValidUser(ctx context.Context, name string, password string) (*models.User, error)
}

type UserService struct {
	CookieHandler *securecookie.SecureCookie
	UserRepo      repository.IUserRepo
}

func NewUserService(cookieHandler *securecookie.SecureCookie, userRepo repository.IUserRepo) IUserService {
	return &UserService{
		CookieHandler: cookieHandler,
		UserRepo:      userRepo,
	}
}

func (s *UserService) GetUserName(request *http.Request) string {
	var cookieValue models.UserCookie

	cookie, err := request.Cookie("session")
	if err != nil {
		return ""
	}

	err = s.CookieHandler.Decode("session", cookie.Value, &cookieValue)
	if err != nil {
		return ""
	}

	return cookieValue.Name
}

func (s *UserService) SetSession(userName string, response http.ResponseWriter) error {
	cookieValue := models.UserCookie{
		Name: userName,
	}

	encoded, err := s.CookieHandler.Encode("session", cookieValue)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}

	http.SetCookie(response, cookie)
	return nil
}

func (s *UserService) ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(response, cookie)
}

func (s *UserService) IsValidUser(ctx context.Context, name string, password string) (*models.User, error) {
	return s.UserRepo.IsValidUser(ctx, name, password)
}
