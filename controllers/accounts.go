package controllers

import (
	"gin-test-example/models/account"
	"gin-test-example/pkg/e"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) (int, interface{}) {
	acc := account.Accounts{}
	err := c.Bind(&acc)
	if err != nil {
		return e.PARAMETER_ERROR, nil
	}
	err = acc.InsertNewAccount()
	if err != nil {
		return e.SERVER_ERROR, err
	}
	return e.SUCCESS, nil
}

func AccountsList(c *gin.Context) (int, interface{})  {
	acc := account.Accounts{}
	accs, err := acc.ListAccounts()
	if err != nil {
		return e.SERVER_ERROR, err
	}
	return e.SUCCESS, accs
}
