package config

import (
	"os"
	"strconv"
)

// Envs
var (
	// SQL
	sqlHost     = getEnv("", "localhost")
	sqlUser     = getEnv("", "root")
	sqlPassword = getEnv("", "password")
	sqlDbName   = getEnv("", "sethr_db")
	sqlPort     = getEnv("", "3306")

	// Redis
	redisHost     = getEnv("", "localhost")
	redisPort     = getEnv("", "6379")
	redisPassword = getEnv("", "")
	redisDb       = getEnvInt("", 0)
	redisTTL      = getEnvInt("", 3600)

	// App
	appHost     = getEnv("", "localhost")
	appPort     = getEnv("", "3000")
	appTokenKey = getEnv("", "tokenKey")
)

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {

		v, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		}

		return v
	}

	return fallback
}
