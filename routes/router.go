package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	RegisterApi(router.Group("api"))
	return router
}
