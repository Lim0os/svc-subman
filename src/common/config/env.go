package config

import "os"

type Env struct {
	Repository Repo
	Server     Server
	LogLvl     string
}

type Repo struct {
	DbPort    string
	DbUser    string
	DbPass    string
	DbName    string
	DbHost    string
	DbSSLMode string
	DbShema   string
}

type Server struct {
	Port string
}

func InitEnv() *Env {
	return &Env{
		LogLvl: getEnv("LOG_LEVEL", "info"),
		Repository: Repo{
			DbPort:    getEnv("DB_PORT", "5432"),
			DbUser:    getEnv("DB_USER", "postgres"),
			DbPass:    getEnv("DB_PASS", "postgres"),
			DbName:    getEnv("DB_NAME", "postgres"),
			DbHost:    getEnv("DB_HOST", "localhost"),
			DbSSLMode: getEnv("DB_SSLMODE", "disable"),
			DbShema:   getEnv("DB_SHEMA", ""),
		},
		Server: Server{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
