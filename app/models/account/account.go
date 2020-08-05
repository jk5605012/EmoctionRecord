package account
//test
type Accounts struct {
	UserName string `json:"username" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
}
