package main

import (
	"fmt"
	"github.com/maximegorov13/go-api/configs"
	"github.com/maximegorov13/go-api/internal/auth"
	"github.com/maximegorov13/go-api/internal/link"
	"github.com/maximegorov13/go-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
