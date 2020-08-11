package account

import (
	"gin-test-example/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestInsert(t *testing.T) {
	fackdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.Nil(t, err)
	db.DB, err = gorm.Open("postgres", fackdb)
	assert.Nil(t, err)
	defer db.DB.Close()
	type fields struct {
		UserName string
		Pwd      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Add one Account",
			fields: fields{
				UserName: "Anish",
				Pwd:      "123456",
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accs := &Accounts{
				UserName: tt.fields.UserName,
				Pwd:      tt.fields.Pwd,
			}
			sqlcmd := `INSERT INTO "accounts" ("user_name","pwd") VALUES ($1,$2) RETURNING "accounts"."id"`
			mock.ExpectBegin()
			mock.ExpectQuery(sqlcmd).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(i))
			mock.ExpectCommit()
			if err := accs.InsertNewAccount(); (err != nil) != tt.wantErr {
				t.Errorf("Accounts.InsertNewAccount() error = %v", err)
			}
		})
	}
}

func TestList(t *testing.T) {
	fackdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.Nil(t, err)
	db.DB, err = gorm.Open("postgres", fackdb)
	assert.Nil(t, err)
	defer db.DB.Close()
	type fields struct {
		UserName string
		Pwd      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "get all accounts",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accs := &Accounts{
				UserName: tt.fields.UserName,
				Pwd:      tt.fields.Pwd,
			}
			sqlCMD := `SELECT * FROM "accounts"`
			mock.ExpectQuery(sqlCMD).WillReturnRows(sqlmock.NewRows(nil))
			if res, err := accs.ListAccounts(); (err != nil) != tt.wantErr {
				t.Error(err)
			} else {
				t.Log(res)
			}
		})
	}
}
