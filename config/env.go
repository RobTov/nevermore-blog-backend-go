package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port                   string
	DatabaseURI            string
	JWTExpirationInSeconds uint64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:                   getEnv("PORT", "8080"),
		DatabaseURI:            getEnv("DATABASE_URI", "postgres://postgres:postgres@127.0.0.1:5497/postgres?sslmode=disable"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "supersecretpassword"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback uint64) uint64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return uint64(i)
	}

	return fallback
}
