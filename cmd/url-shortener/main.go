package main

import (
	"fmt"

	"github.com/Paheter1/url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
