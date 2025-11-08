package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	FirebaseConfigPath string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	cfg := &Config{
		Port:               getEnv("PORT", "8080"), 
		FirebaseConfigPath: getEnv("FIREBASE_CONFIG_PATH", "/etc/secrets/serviceAccountKey.json"),
	}

	// Helpful log for debugging deployment
	log.Printf("✅ Loaded config: PORT=%s, FIREBASE_CONFIG_PATH=%s", cfg.Port, cfg.FirebaseConfigPath)

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
