package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	FirebaseConfigPath string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	// fmt.Println(err);
	// if err != nil {
	// 	log.Println("Error loading .env file, using system environment variables")
	// }

	return &Config{
		Port:               getEnv("PORT", "8080"),
		FirebaseConfigPath: getEnv("FIREBASE_CONFIG_PATH", "/etc/secrets/serviceAccountKey.json"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
