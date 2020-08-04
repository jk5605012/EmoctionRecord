package main

import (
	"gin-test-example/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.Init()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
