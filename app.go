package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erhemdiputra/go-di/config"
	"github.com/erhemdiputra/go-di/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("[Err] Load Config File: %+v", err)
	}

	router := httprouter.New()

	userHandler := handler.NewUserHandler(router)
	userHandler.Serve()

	port := config.Main.Server.Port
	log.Printf("Listening on Port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
