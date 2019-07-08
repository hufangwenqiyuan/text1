package router

import (
	"github.com/gin-gonic/gin"
	. "text1/text1/src/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)

	router.POST("/user", Store)

	router.PUT("/user/:id", Update)

	router.DELETE("user/:id", Destroy)

	return router
}
