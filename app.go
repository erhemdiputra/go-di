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
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()

	handler.NewPlayerHandler(router, database.Get(), infraMemCache.GetKodingCache())
	handler.NewUserHandler(router)

	http.Handle("/", router)

	port := globalCfg.Server.Port
	log.Printf("Listening on Port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
