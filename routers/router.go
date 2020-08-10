package routers

import (
	"gin-test-example/controllers"
	"gin-test-example/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) (int, interface{})

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		errcode, data := handler(c)
		switch errcode {
		case e.PARAMETER_ERROR:
			controllers.Response(c, http.StatusBadRequest, e.PARAMETER_ERROR, data)
		case e.SERVER_ERROR:
			controllers.Response(c, http.StatusInternalServerError, e.SERVER_ERROR, data)
		default:
			controllers.Response(c, http.StatusOK, e.SUCCESS, data)
		}
	}
}

func Init() *gin.Engine {
	r := gin.Default()
	r.POST("/register", wrapper(controllers.Register))
	// r.GET("/accounts", )

	return r
}
