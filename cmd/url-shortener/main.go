package main

import (
	"fmt"

	"github.com/al-grigorian/url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: "chi render"

	// TODO: run server
}
