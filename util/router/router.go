package router

import (
	"github.com/gin-gonic/gin"
	. "happyD/util/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)
	router.GET("/odds", Odds)

	router.POST("/user", Store)
	router.POST("/oddp", Oddp)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)

	return router
}
