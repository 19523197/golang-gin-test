package router

import (
	_ "belajar-gin/handler"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Ping)
		product := v1.Group("/product")
		{
			product.GET("/", nil)
			product.GET("/:id", nil)
		}
	}
	return r
}

func Ping(c *gin.Context) {
	c.String(200, "ping")
}
