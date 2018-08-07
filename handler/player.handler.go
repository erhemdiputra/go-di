package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	infraMemCache "github.com/erhemdiputra/go-di/infrastructure_services/memcache"
	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/repository"
	"github.com/erhemdiputra/go-di/service"
)

type PlayerHandler struct {
	PlayerService service.IPlayerService
}

func NewPlayerHandler(db *sql.DB, memCache *infraMemCache.KodingCache) *PlayerHandler {
	playerRepo := repository.NewPlayerRepo(db)
	playerService := service.NewPlayerService(playerRepo, memCache)

	return &PlayerHandler{
		PlayerService: playerService,
	}
}

func (h *PlayerHandler) Serve() {
	http.Handle("/api/player/list", HandlerFunc(h.GetList))
	http.Handle("/api/player/add", HandlerFunc(h.Add))
	http.Handle("/api/player/", HandlerFunc(h.GetByID))
	http.Handle("/api/player/update/", HandlerFunc(h.Update))
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

	list, err := h.PlayerService.GetList(ctx, form)
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

	form.Sanitize()

	_, err = h.PlayerService.Add(ctx, form)
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
	if err != nil || id <= 0 {
		return nil, errors.New("Invalid player ID")
	}

	player, err := h.PlayerService.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	return player, nil
}

func (h *PlayerHandler) Update(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("Invalid Request")
	}

	ctx := r.Context()
	strID := r.URL.Path[len("/api/player/update/"):]

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil || id <= 0 {
		return nil, errors.New("Invalid Player ID")
	}

	var form models.PlayerForm
	err = GetJSONParams(r, &form)
	if err != nil || form.IsEmpty() {
		return nil, errors.New("Invalid JSON Params")
	}

	_, err = h.PlayerService.Update(ctx, id, form)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	resp := ResponseStatus{
		IsSuccess: 1,
	}

	return resp, nil
}
