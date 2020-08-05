package routers

import (
	"gin-test-example/controllers"
	"gin-test-example/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) int

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		errcode := handler(c)
		switch errcode {
		case e.PARAMETER_ERROR:
			controllers.Response(c, http.StatusBadRequest, e.PARAMETER_ERROR, nil)
		default:
			controllers.Response(c, http.StatusOK, e.SUCCESS, nil)
		}
	}
}

func Init() *gin.Engine {
	r := gin.Default()
	r.POST("/register", wrapper(controllers.Register))

	return r
}
