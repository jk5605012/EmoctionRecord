package integration

import (
	"gin-test-example/db"
	"gin-test-example/models/account"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	var err error
	db.DB, err = gorm.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestAccounts_InsertNewAccount(t *testing.T) {
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
		{
			name: "Add same Account",
			fields: fields{
				UserName: "Anish",
				Pwd:      "123456",
			},
			wantErr: true,
		},
	}
	db.DB.CreateTable(&account.Accounts{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accs := &account.Accounts{
				UserName: tt.fields.UserName,
				Pwd:      tt.fields.Pwd,
			}
			if err := accs.InsertNewAccount(); (err != nil) != tt.wantErr {
				t.Errorf("Accounts.InsertNewAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	db.DB.DropTable(&account.Accounts{})

}
