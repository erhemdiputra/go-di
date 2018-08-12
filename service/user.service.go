package service

import (
	"net/http"

	"github.com/erhemdiputra/go-di/models"
	"github.com/gorilla/securecookie"
)

type IUserService interface {
	GetUserName(request *http.Request) string
	SetSession(userName string, response http.ResponseWriter) error
	ClearSession(response http.ResponseWriter)
}

type UserService struct {
	CookieHandler *securecookie.SecureCookie
}

func NewUserService(cookieHandler *securecookie.SecureCookie) IUserService {
	return &UserService{
		CookieHandler: cookieHandler,
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
