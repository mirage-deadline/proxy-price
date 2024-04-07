package config

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Config struct {
	RedisDSN      string
	RedisPassword string
	BinanceURL    string
	Symbols       []string
	Topic         string
}

func MustGetConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Ошибка при загрузке .env файла: " + err.Error())
	}

	symbols := parseSymbols(getEnv("SYMBOLS"))

	return &Config{
		RedisDSN:      getEnv("REDIS_DSN"),
		RedisPassword: getEnv("REDIS_PASSWORD"),
		BinanceURL:    getEnv("BINANCE_URL"),
		Symbols:       symbols,
		Topic:         getEnv("TOPIC"),
	}
}

func parseSymbols(symbols string) []string {
	return strings.Split(symbols, ", ")
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return ""
	}
	return value
}
