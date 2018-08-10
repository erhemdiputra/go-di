package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/erhemdiputra/go-di/config"
	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Status            string      `json:"status"`
	ServerProcessTime string      `json:"server_process_time"`
	Data              interface{} `json:"data"`
	MessageError      []string    `json:"message_error,omitempty"`
}

type ResponseStatus struct {
	IsSuccess int `json:"is_success"`
}

func wrapHandler(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()

		timeout, err := time.ParseDuration(config.Get().Server.Timeout)
		if err != nil {
			timeout = 5 * time.Second
		}

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		r = r.WithContext(ctx)
		fn(w, r, ps)
	}
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoded, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.Write(encoded)
	return
}

func writeSuccess(w http.ResponseWriter, processTime float64, data interface{}) {
	resp := Response{
		Status:            http.StatusText(http.StatusOK),
		ServerProcessTime: fmt.Sprintf("%f", processTime),
		Data:              data,
	}

	writeJSONResponse(w, http.StatusOK, resp)
	return
}

func writeError(w http.ResponseWriter, processTime float64, message ...string) {
	resp := Response{
		Status:            http.StatusText(http.StatusOK),
		ServerProcessTime: fmt.Sprintf("%f", processTime),
		MessageError:      message,
		Data:              ResponseStatus{},
	}

	writeJSONResponse(w, http.StatusOK, resp)
	return
}
