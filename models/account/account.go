package account

type Accounts struct {
	UserName string `json:"username" gorm:"unique" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
}
