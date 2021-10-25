package configs

import (
	"log"
	"os"
)

type Config struct {
	SecretHashKey string
}

var config Config

func LoadConfig() error {

	if key := os.Getenv("SECRET_HASH_KEY"); key == "" {
		log.Fatalf("no se ha especificado SECRET_HASH_KEY")
		config.SecretHashKey = key
	}
	return nil
}


func GetConfig() *Config {
	return &config
}