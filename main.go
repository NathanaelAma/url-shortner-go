package main

import (
	"github.com/NathanaelAma/url-shortener/internal"
	"os"

	"github.com/NathanaelAma/url-shortener/config"
)

func main() {
	config.Runtime()
	internal.StartServer("./", os.Getenv("PORT"))
}
