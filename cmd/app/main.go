package main

import (
	"log"

	"github.com/Ayeye11/se-thr/config"
	"github.com/Ayeye11/se-thr/infrastructure/api"
	"github.com/Ayeye11/se-thr/infrastructure/database"
	"github.com/Ayeye11/se-thr/infrastructure/server"
)

func main() {
	// Config
	cfg := config.LoadConfig()

	// Database
	sqlDB, err := database.InitSQL("mysql", cfg.SQL, 3)
	if err != nil {
		log.Fatalf("fatal: %v\n", err)
	}

	// API Internals
	router := api.NewRouter(sqlDB.GetDB(), cfg.APP.TokenKey)
	handler := router.RegisterRoutes()

	// Server
	server := server.NewServer(handler, cfg.APP)

	// server:Run
	go server.Run()

	// server:Close
	if err := server.Close(); err != nil {
		log.Fatalf("fatal: %v\n", err)
	}

	// database:Close
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("fatal: %v\n", err)
	}

	log.Println("API Closed gracefully")
}
