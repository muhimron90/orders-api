package config

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() Config {
	cfg := Config{

		RedisAddress: "127.0.0.1:6379",
		ServerPort:   3000,
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDRESS"); exists {
		cfg.RedisAddress = redisAddr
	}
	if portNumber, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseInt(portNumber, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg
}
