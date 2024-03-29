package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World")
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", handler)

	fmt.Println("Go REST service started...")
	router.Run("0.0.0.0:9009")
}

