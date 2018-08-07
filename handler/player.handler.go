package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/erhemdiputra/go-di/controller"
	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/repository"
	"github.com/erhemdiputra/go-di/service"
)

type PlayerHandler struct {
	PlayerController *controller.PlayerController
}

func NewPlayerHandler(db *sql.DB) *PlayerHandler {
	playerRepo := repository.NewPlayerRepo(db)
	playerService := service.NewPlayerService(playerRepo)
	playerController := controller.NewPlayerController(playerService)

	return &PlayerHandler{
		PlayerController: playerController,
	}
}

func (h *PlayerHandler) Serve() {
	http.Handle("/api/player/list", HandlerFunc(h.GetList))
	http.Handle("/api/player/add", HandlerFunc(h.Add))
	http.Handle("/api/player/", HandlerFunc(h.GetByID))
}

func (h *PlayerHandler) GetList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("Invalid Request")
	}

	ctx := r.Context()

	var form models.PlayerForm
	if err := GetJSONParams(r, &form); err != nil {
		return nil, errors.New("Invalid JSON Params")
	}

	list, err := h.PlayerController.GetList(ctx, form)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	return list, nil
}

func (h *PlayerHandler) Add(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("Invalid Request")
	}

	ctx := r.Context()

	var form models.PlayerForm
	err := GetJSONParams(r, &form)
	if err != nil || form.IsEmpty() {
		return nil, errors.New("Invalid JSON Params")
	}

	_, err = h.PlayerController.Add(ctx, form)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	resp := ResponseStatus{
		IsSuccess: 1,
	}

	return resp, nil
}

func (h *PlayerHandler) GetByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != http.MethodGet {
		return nil, errors.New("Invalid Request")
	}

	ctx := r.Context()
	strID := r.URL.Path[len("/api/player/"):]

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		return nil, errors.New("Invalid player ID")
	}

	player, err := h.PlayerController.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	return player, nil
}
