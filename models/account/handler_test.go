package account

import (
	"fmt"
	"gin-test-example/db"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ory/dockertest"
)

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
	}
	pool, err := dockertest.NewPool("tcp://localhost:2375")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := pool.Run("postgres", "9.6", []string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=test"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	if err = resource.Expire(60); err != nil {
		log.Fatal(err)
	}
	err = pool.Retry(func() error {
		var err error
		db.DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), "test"))
		if err != nil {
			return err
		}
		return db.DB.DB().Ping()
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	db.DB.CreateTable(&Accounts{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accs := &Accounts{
				UserName: tt.fields.UserName,
				Pwd:      tt.fields.Pwd,
			}
			if err := accs.InsertNewAccount(); (err != nil) != tt.wantErr {
				t.Errorf("Accounts.InsertNewAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err = pool.Purge(resource); err != nil {
		log.Fatal(err)
	}
}
