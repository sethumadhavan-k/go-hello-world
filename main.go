package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

func handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World")
}

func printTime(ctx *gin.Context) {
	t := time.Now()
	ctx.JSON(http.StatusOK, t.Format("20060102150405"))
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", handler)
	router.GET("/time", printTime)

	fmt.Println("Go REST service started...")
	router.Run("0.0.0.0:9009")
}

