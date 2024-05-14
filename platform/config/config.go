package config

import (
	"os"
	"strconv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBHost                 string
	DBPort                 string
	DBName                 string
	JWTExpirationInSeconds int
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "127.0.0.1"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "root"),
		DBHost:                 getEnv("DB_HOST", "127.0.0.1"),
		DBPort:                 getEnv("DB_PORT", "3306"),
		DBName:                 getEnv("DB_NAME", "concursos-api"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600),
		JWTSecret:              getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
