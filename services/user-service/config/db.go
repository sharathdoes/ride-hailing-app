package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB(cfg *Config) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("MongoDB ping failed:", err)
    }

    DB = client.Database(cfg.MongoDB)
    log.Println("MongoDB connected")
}