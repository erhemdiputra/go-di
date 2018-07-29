package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/erhemdiputra/go-di/config"
)

type Response struct {
	Data interface{} `json:"data"`
}

type ResponseStatus struct {
	IsSuccess int `json:"is_success"`
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	timeout, err := time.ParseDuration(config.Get().Server.Timeout)
	if err != nil {
		timeout = time.Second * 5
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	r = r.WithContext(ctx)

	res, err := fn(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := Response{
		Data: res,
	}

	bytesResp, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesResp)
}
