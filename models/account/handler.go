package account

import "gin-test-example/db"

func (accs *Accounts) InsertNewAccount() (err error) {
	err = db.DB.Create(accs).Error
	return err
}
