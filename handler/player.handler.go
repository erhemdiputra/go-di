package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	infraMemCache "github.com/erhemdiputra/go-di/infrastructure_services/memcache"
	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/repository"
	"github.com/erhemdiputra/go-di/service"
	"github.com/julienschmidt/httprouter"
)

type PlayerHandler struct {
	PlayerService service.IPlayerService
}

func NewPlayerHandler(router *httprouter.Router, db *sql.DB, memCache *infraMemCache.KodingCache) {
	playerRepo := repository.NewPlayerRepo(db)
	playerService := service.NewPlayerService(playerRepo, memCache)

	handler := &PlayerHandler{
		PlayerService: playerService,
	}

	router.POST("/api/player/list", wrapHandler(handler.GetList))
	router.POST("/api/player/add", wrapHandler(handler.Add))
	router.GET("/api/player/:id", wrapHandler(handler.GetByID))
	router.POST("/api/player/update/:id", wrapHandler(handler.Update))
}

func (h *PlayerHandler) GetList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	ctx := r.Context()

	var form models.PlayerForm
	if err := GetJSONParams(r, &form); err != nil {
		writeError(w, time.Since(start).Seconds(), "Invalid JSON Params")
		return
	}

	list, err := h.PlayerService.GetList(ctx, form)
	if err != nil {
		writeError(w, time.Since(start).Seconds(), "Internal Server Error")
		return
	}

	writeSuccess(w, time.Since(start).Seconds(), list)
}

func (h *PlayerHandler) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	ctx := r.Context()

	var form models.PlayerForm
	err := GetJSONParams(r, &form)
	if err != nil || form.IsEmpty() {
		writeError(w, time.Since(start).Seconds(), "Invalid JSON Params")
		return
	}

	form.Sanitize()

	_, err = h.PlayerService.Add(ctx, form)
	if err != nil {
		writeError(w, time.Since(start).Seconds(), "Internal Server Error")
		return
	}

	resp := ResponseStatus{
		IsSuccess: 1,
	}

	writeSuccess(w, time.Since(start).Seconds(), resp)
}

func (h *PlayerHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	ctx := r.Context()
	idStr := ps.ByName("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		writeError(w, time.Since(start).Seconds(), "Invalid Player ID")
		return
	}

	player, err := h.PlayerService.GetByID(ctx, id)
	if err != nil {
		writeError(w, time.Since(start).Seconds(), "Internal Server Error")
		return
	}

	writeSuccess(w, time.Since(start).Seconds(), player)
}

func (h *PlayerHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	ctx := r.Context()
	idStr := ps.ByName("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		writeError(w, time.Since(start).Seconds(), "Invalid Player ID")
		return
	}

	var form models.PlayerForm
	err = GetJSONParams(r, &form)
	if err != nil || form.IsEmpty() {
		writeError(w, time.Since(start).Seconds(), "Invalid JSON Params")
		return
	}

	_, err = h.PlayerService.Update(ctx, id, form)
	if err != nil {
		writeError(w, time.Since(start).Seconds(), "Internal Server Error")
		return
	}

	resp := ResponseStatus{
		IsSuccess: 1,
	}

	writeSuccess(w, time.Since(start).Seconds(), resp)
}
