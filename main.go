package main

import (
    "github.com/gin-gonic/gin"
    "go-auth-api/handlers"
    "go-auth-api/middleware"
	"net/http"
)
func initDatabase() {
    var err error
    models.DB, err = gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    fmt.Println("Database connected!")

    // Auto-migrate the User model
    models.DB.AutoMigrate(&models.User{})
}
func main() {
    r := gin.Default()

    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    // Protected route to fetch user details
    r.GET("/user", middleware.AuthMiddleware(), func(c *gin.Context) {
        username := c.MustGet("username").(string)
        c.JSON(http.StatusOK, gin.H{"username": username})
    })

    r.Run(":8080") // Start the server on port 8080
}
