package main

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

func PostProducts(ctx	*gin.Context) {
	var newProduct Product

	if err := ctx.BindJSON(&newProduct); err != nil {
		return
	}
	products = append(products, newProduct)
	ctx.JSON(http.StatusCreated, products)
}

func GetProductById(ctx *gin.Context) {
	for _, a := range products {
		if a.ID == ctx.Param("id") {
			ctx.JSON(http.StatusOK, a)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "the Product you are looking for is not found"})
}

func UpdateProductById(ctx *gin.Context) {
	var product Product
	if err := ctx.BindJSON(&product); err != nil {
		return
	}
	for i, a := range products {
		if a.ID == ctx.Param("id") {
			products[i].ID = product.ID
			products[i].Name = product.Name 
			products[i].Price = product.Price 
			products[i].Stock = product.Stock
			ctx.JSON(http.StatusAccepted, products[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "can't find the Product id you are searching for"})
}

func DeleteProductById(ctx *gin.Context) {
	var product Product
	if err := ctx.BindJSON(&product); err != nil {
		return
	}
	for i, u := range products {
		if u.ID == ctx.Param("id") {
			products = append(products[:i], products[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Product Deleted successfuly"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

