package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rejvban/macro-api/internal/router"
)

func main() {
	fmt.Println("Hello World")

	router := router.New()

	server := &http.Server{
		Addr:         "0.0.0.0:3000",
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	fmt.Println("Starting server on port 3000")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
