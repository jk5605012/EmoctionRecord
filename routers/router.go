package routers

import (
	"gin-test-example/controllers"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)

	return r
}
