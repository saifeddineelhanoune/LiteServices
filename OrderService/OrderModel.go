package main

import (
	"net/http"
)

// Request to create an order
type CreateOrderRequest struct {
    UserID string              `json:"user_id"`
    Items  []CreateOrderItem   `json:"items"`
}

type CreateOrderItem struct {
    ProductID string `json:"product_id"`
    Quantity  int    `json:"quantity"`
}

type Order struct {
    ID        string      `json:"id"`
    UserID    string      `json:"user_id"`
    Items     []OrderItem `json:"items"`
    Total     float64     `json:"total"`
    Status    string      `json:"status"`
}

type OrderItem struct {
    ProductID   string  `json:"product_id"`
    ProductName string  `json:"product_name"`
    UnitPrice   float64 `json:"unit_price"`
    Quantity    int     `json:"quantity"`
}

