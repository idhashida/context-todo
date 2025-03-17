package main

import (
	"fmt"
	"go/context-todo/configs"
	"go/context-todo/internal/auth"
	"net/http"
)

func main() {
	_ = configs.LoadConfig()
	router := http.NewServeMux()

	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("[ START ] server is listening on port 8081")
	server.ListenAndServe()
	server.ListenAndServe()
}
