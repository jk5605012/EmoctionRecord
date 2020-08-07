package account

type Accounts struct {
	ID       int    `json:"-" gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserName string `json:"user_name" gorm:"unique" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
}
