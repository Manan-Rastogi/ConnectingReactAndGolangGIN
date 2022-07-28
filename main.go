package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
)

func main() {
	server := gin.Default()

	server.Use(gin.Recovery(), gin.Logger(), CORSMiddleware())
	

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status" : "working",
		})
	})

	server.POST("/sum", func(ctx *gin.Context) {
		cookie := ctx.Request.Header.Get("cookies")
		a := 10
		ctx.JSON(200, gin.H{
			"yourcookie":cookie,
			"sum" : a + 12,
		})
	})

	server.Run(":8009")
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}