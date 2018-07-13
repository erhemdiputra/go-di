package main

import (
	"log"
	"net/http"

	"github.com/erhemdiputra/practice-mvc/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	userHandler := handler.NewUserHandler(router)
	userHandler.Serve()

	log.Println("Listening on Port 8080")
	http.ListenAndServe(":8080", router)
}
