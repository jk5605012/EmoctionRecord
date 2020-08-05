package account

import "github.com/jinzhu/gorm"

type Accounts struct {
	gorm.Model
	UserName string `json:"username" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
}
