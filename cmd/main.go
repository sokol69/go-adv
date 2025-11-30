package main

import (
	"adv-demo/configs"
	"adv-demo/internal/auth"
	"adv-demo/internal/link"
	"adv-demo/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server on port 8081")
	server.ListenAndServe()
}
