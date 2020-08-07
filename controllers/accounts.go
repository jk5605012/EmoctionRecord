package controllers

import (
	"gin-test-example/models/account"
	"gin-test-example/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) (int, interface{}) {
	acc := account.Accounts{}
	err := c.Bind(&acc)
	if err != nil {
		return e.PARAMETER_ERROR, nil
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
	return e.SUCCESS, nil
}
