package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name       string  `json:"name"`
    Price      float64 `json:"price"`
    Quantity   uint    `json:"quantity"`
    ImageURL   string  `json:"image_url"`
}
