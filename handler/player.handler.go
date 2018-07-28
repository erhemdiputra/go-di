package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/erhemdiputra/go-di/controller"
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
}

func (h *PlayerHandler) GetList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != http.MethodGet {
		return nil, errors.New("Invalid Request")
	}

	ctx := r.Context()
	return h.PlayerController.GetList(ctx)
}
