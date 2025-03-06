package config

import "os"

// Envs
var (
	// SQL
	sqlHost     = getEnv("", "localhost")
	sqlUser     = getEnv("", "root")
	sqlPassword = getEnv("", "password")
	sqlDbName   = getEnv("", "test_db")
	sqlPort     = getEnv("", "3306")

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
