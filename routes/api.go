package routes

import (
	"gin-restful-demo/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterApi(api *gin.RouterGroup) {
	api.POST("/coach", controllers.InsertCoach)
}
