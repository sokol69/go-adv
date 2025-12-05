package main

import (
	"adv-demo/configs"
	"adv-demo/internal/auth"
	"adv-demo/internal/link"
	"adv-demo/pkg/db"
	"adv-demo/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(database)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server on port 8081")
	server.ListenAndServe()
}
