package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port        string
    MongoURI    string
    MongoDB     string
    KafkaBroker string
    JWTSecret   string
}

func Load() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file, reading from environment")
    }

    return &Config{
        Port:        getEnv("PORT", "8081"),
        MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
        MongoDB:     getEnv("MONGO_DB", "rideshare"),
        KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
        JWTSecret:   getEnv("JWT_SECRET", "changeme"),
    }
}

func getEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}