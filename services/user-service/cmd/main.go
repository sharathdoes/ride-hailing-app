package main


import (
    "rideshare/user-service/config"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    cfg := config.Load()
    config.ConnectDB(cfg)

    r := gin.Default()
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "user-service ok"})
    })

    log.Fatal(r.Run(":" + cfg.Port))
}


/*
POST   /auth/register        → register as rider or driver
POST   /auth/login           → returns JWT token
POST   /auth/logout          → invalidate token (optional)
POST   /auth/refresh         → refresh JWT token

GET    /users/me             → get my profile
PUT    /users/me             → update my profile (name, phone, photo)

GET    /drivers/me           → get driver profile
PUT    /drivers/me           → update driver details (vehicle info, availability)
PUT    /drivers/me/status    → toggle online/offline
*/