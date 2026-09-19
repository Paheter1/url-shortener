package main

import (
	"fmt"
	"log"

	"github.com/Paheter1/url-shortener/internal/config"
	httpserver "github.com/Paheter1/url-shortener/internal/http-server"
)

func main() {
	cfg := config.MustLoad()
	server := httpserver.NewServer(
		cfg.Address,
		cfg.Timeout,
		cfg.IdleTimeout,
	)
	log.Println("Server Started!")

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
