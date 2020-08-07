package controllers_test

import (
	"bytes"
	"encoding/json"
	"gin-test-example/models/account"
	"gin-test-example/routers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	r := routers.Init()
	w := httptest.NewRecorder()
	body := account.Accounts{
		UserName: "Anish",
		Pwd:      "12345",
	}
	jsonbyte, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonbyte))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	defer w.Result().Body.Close()
	bodyl, _ := ioutil.ReadAll(w.Result().Body)
	t.Log(string(bodyl))
	assert.Equal(t, http.StatusOK, w.Code)
}
