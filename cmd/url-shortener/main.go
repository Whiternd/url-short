package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanenv библиотека
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage инициализируем сторадже: sqllite

	// TODO: init router: chi, render

	// TODO: run server

}
