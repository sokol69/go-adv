package main

import (
	"adv-demo/configs"
	"adv-demo/internal/auth"
	"adv-demo/internal/link"
	"adv-demo/internal/user"
	"adv-demo/pkg/db"
	"adv-demo/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)

	// Services
	authService := auth.NewAuthService(userRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr: ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server on port 8081")
	server.ListenAndServe()
}
