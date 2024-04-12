package handlers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/Roy009/go-shopping-cart/backend/models"
)

// Handler to get all products
func GetProducts(c *gin.Context) {
    var products []models.Product
    db.Find(&products)
    c.JSON(200, products)
}

// Handler to create a new product
func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    db.Create(&product)
    c.JSON(201, product)
}
