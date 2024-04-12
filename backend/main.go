package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/Roy009/go-shopping-cart/backend/models"
    "github.com/Roy009/go-shopping-cart/backend/handlers"
)

var db *gorm.DB

func main() {
    // Initialize Gin router
    r := gin.Default()

    // Initialize GORM
    var err error
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&models.Product{})

    // Routes
    r.GET("/products", handlers.GetProducts)
    r.POST("/products", handlers.CreateProduct)

    // Run the server
    r.Run(":8080")
}
