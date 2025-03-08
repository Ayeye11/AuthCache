package main

import (
	"log"

	"github.com/Ayeye11/se-thr/config"
	"github.com/Ayeye11/se-thr/infrastructure/api"
	"github.com/Ayeye11/se-thr/infrastructure/redis"
	"github.com/Ayeye11/se-thr/infrastructure/server"
	"github.com/Ayeye11/se-thr/infrastructure/sql"
)

func main() {
	// Config
	cfg := config.LoadConfig()

	// Database
	sqlDB, err := sql.InitSQL("mysql", cfg.SQL, 3)
	if err != nil {
		log.Fatalf("fatal: %v\n", err)
	}

	// Redis
	rdb, err := redis.NewRedisDB(cfg.Redis)
	if err != nil {
		log.Fatalf("fatal: %v\n", err)
	}

	// API Internals
	router := api.NewRouter(sqlDB.GetDB(), rdb.GetClient(), rdb.GetTTL(), cfg.APP.TokenKey)
	handler := router.RegisterRoutes()

	// Server
	server := server.NewServer(handler, cfg.APP)

	// server:Run
	go server.Run()

	// server:Close
	if err := server.Close(); err != nil {
		log.Fatalf("fatal: %v\n", err)
	}

	// redis:Close
	if err := rdb.Close(); err != nil {
		log.Fatalf("fatal: %v\n", err)
	}

	// database:Close
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("fatal: %v\n", err)
	}

	log.Println("API Closed gracefully")
}
