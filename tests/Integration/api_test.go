package integration

import (
	"bytes"
	"encoding/json"
	"gin-test-example/db"
	"gin-test-example/models/account"
	"gin-test-example/routers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestRegister(t *testing.T) {
	r := routers.Init()
	tests := []struct {
		name     string
		body     account.Accounts
		wantCode int
	}{
		{
			name: "First Account",
			body: account.Accounts{
				UserName: "Anish",
				Pwd:      "12345",
			},
			wantCode: http.StatusOK,
		},
	}
	db.DB.CreateTable(&account.Accounts{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			jsonbyte, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonbyte))
			req.Header.Add("Content-Type", "application/json")
			r.ServeHTTP(w, req)
			defer w.Result().Body.Close()
			bodyl, _ := ioutil.ReadAll(w.Result().Body)
			t.Log(string(bodyl))
			assert.Equal(t, http.StatusOK, w.Code)
		})
	}
	db.DB.DropTable(&account.Accounts{})
}
