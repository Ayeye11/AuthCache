package config

import "os"

// Envs
var (
	// SQL
	sqlHost     = getEnv("", "localhost")
	sqlUser     = getEnv("", "root")
	sqlPassword = getEnv("", "password")
	sqlDbName   = getEnv("", "thr_sql")
	sqlPort     = getEnv("", ":3306")

	// App
	appHost     = getEnv("", "localhost")
	appPort     = getEnv("", ":8080")
	appTokenKey = getEnv("", "tokenKey")
)

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}
