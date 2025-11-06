package ProductService

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
    ID    string  `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    Stock int     `json:"stock"`
}

var products = []Product {
	{ID: "1", Name: "shampoo", Price: 19.99, Stock: 20},
	{ID: "2", Name: "yagourt", Price: 2.50, Stock: 6},
	{ID: "3", Name: "candle", Price: 3.00, Stock: 40},
}

func GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, products)
}

