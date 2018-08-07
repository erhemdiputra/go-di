package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erhemdiputra/go-di/config"
	"github.com/erhemdiputra/go-di/database"
	"github.com/erhemdiputra/go-di/handler"
	infraMemCache "github.com/erhemdiputra/go-di/infrastructure_services/memcache"
	_ "github.com/go-sql-driver/mysql"
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

	infraMemCache.InitKodingCache()

	playerHandler := handler.NewPlayerHandler(database.Get(), infraMemCache.GetKodingCache())
	playerHandler.Serve()

	port := globalCfg.Server.Port
	log.Printf("Listening on Port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
