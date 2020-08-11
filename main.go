package main

import (
	"gin-test-example/db"
	"gin-test-example/models/account"
	"gin-test-example/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.Init()
	db.Init()
	db.DB.AutoMigrate(&account.Accounts{})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
