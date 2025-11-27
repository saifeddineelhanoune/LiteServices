package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	UserID string            `json:"user_id"`
	Items  []CreateOrderItem `json:"items"`
}

type CreateOrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	UserName  string      `json:"user_name"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt string      `json:"created_at"`
}

type OrderItem struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var orders = []Order{}
var orderCounter = 1

const (
	UserServiceURL    = "http://user-service:5000"
	ProductServiceURL = "http://product-service:5001"
)

func main() {
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Order Service Running"})
	})
	route.POST("/orders", CreateOrder)
	route.GET("/orders", GetOrders)
	route.GET("/orders/:id", GetOrderById)

	route.Run("0.0.0.0:5002")
}

func CreateOrder(ctx *gin.Context) {
	var req CreateOrderRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.UserID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	if len(req.Items) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "at least one item is required"})
		return
	}

	user, err := getUserById(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %v", err)})
		return
	}

	var orderItems []OrderItem
	var total float64

	for _, item := range req.Items {
		if item.Quantity <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid quantity for product %s", item.ProductID)})
			return
		}

		product, err := getProductById(item.ProductID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Product %s not found: %v", item.ProductID, err)})
			return
		}

		if product.Stock < item.Quantity {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Insufficient stock for product %s. Available: %d, Requested: %d",
					product.Name, product.Stock, item.Quantity),
			})
			return
		}

		subtotal := product.Price * float64(item.Quantity)
		orderItem := OrderItem{
			ProductID:   product.ID,
			ProductName: product.Name,
			UnitPrice:   product.Price,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		}

		orderItems = append(orderItems, orderItem)
		total += subtotal
	}

	order := Order{
		ID:        strconv.Itoa(orderCounter),
		UserID:    user.ID,
		UserName:  user.Name,
		Items:     orderItems,
		Total:     total,
		Status:    "pending",
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	orderCounter++
	orders = append(orders, order)

	ctx.JSON(http.StatusCreated, order)
}

func GetOrders(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, orders)
}

func GetOrderById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, order := range orders {
		if order.ID == id {
			ctx.JSON(http.StatusOK, order)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
}

func getUserById(userID string) (*User, error) {
	url := fmt.Sprintf("%s/users/%s", UserServiceURL, userID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to User Service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("user service returned status %d: %s", resp.StatusCode, string(body))
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user response: %v", err)
	}

	return &user, nil
}

func getProductById(productID string) (*Product, error) {
	url := fmt.Sprintf("%s/products/%s", ProductServiceURL, productID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Product Service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("product service returned status %d: %s", resp.StatusCode, string(body))
	}

	var product Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, fmt.Errorf("failed to decode product response: %v", err)
	}

	return &product, nil
}
