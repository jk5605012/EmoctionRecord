package account

import "gin-test-example/db"

func (accs *Accounts) InsertNewAccount() (err error) {
	err = db.DB.Create(accs).Error
	return err
}

func (accs *Accounts) ListAccounts() (res []Accounts, err error) {
	err = db.DB.Find(&res).Error
	return res, err
}