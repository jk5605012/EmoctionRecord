package controllers

import (
	"gin-test-example/pkg/e"

	"github.com/gin-gonic/gin"
)

func Response(g *gin.Context, httpCode, errCode int, data interface{}) {
	g.JSON(httpCode, gin.H{
		"http_status": httpCode,
		"code":        errCode,
		"msg":         e.GetMsg(errCode),
		"data":        data,
	})
}
