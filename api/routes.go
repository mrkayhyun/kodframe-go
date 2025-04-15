package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kodframe-go/api/handlers"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", handlers.HelloworldHandler)
		}
	}
}
