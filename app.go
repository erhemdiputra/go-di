package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erhemdiputra/go-di/config"
	"github.com/erhemdiputra/go-di/database"
	"github.com/erhemdiputra/go-di/handler"
	infraMemCache "github.com/erhemdiputra/go-di/infrastructure_services/memcache"
	"github.com/erhemdiputra/go-di/views"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/securecookie"
	"github.com/julienschmidt/httprouter"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("[ERR] Initiate config error: %+v", err)
	}

	globalCfg := config.Get()

	driver, connString := globalCfg.Database.Driver, globalCfg.Database.String()
	if err := database.Init(driver, connString); err != nil {
		log.Fatalf("[ERR] Initiate database error: %+v", err)
	}
	defer database.Get().Close()

	views.PopulateTemplate()
	infraMemCache.InitKodingCache()
	router := httprouter.New()

	cookieHandler := securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	)

	handler.NewPlayerHandler(router, database.Get(), infraMemCache.GetKodingCache())
	handler.NewUserHandler(router, database.Get(), cookieHandler, views.GetMapTemplate())

	port := globalCfg.Server.Port
	log.Printf("Listening on Port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
