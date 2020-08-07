package account

type Accounts struct {
	ID       int    `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserName string `json:"username" gorm:"unique" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
}
