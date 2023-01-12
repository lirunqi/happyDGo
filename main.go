package main

import (
	"github.com/gin-gonic/gin"
	_ "happyD/util/database"
	orm "happyD/util/database"
	"happyD/util/middlewares"
	"happyD/util/router"
)

func main() {
	defer orm.Eloquent.Close()
	r := gin.Default()
	r.Use(middlewares.Cors())
	router := router.InitRouter()
	router.Run(":8000")
}
