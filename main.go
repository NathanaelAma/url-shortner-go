package main

import (
	"os"

	"github.com/NathanaelAma/url-shortener/internal"

	"github.com/NathanaelAma/url-shortener/config"
)

func main() {
	config.Runtime()
	internal.StartServer("./", os.Getenv("PORT"), os.Getenv("ENV"))
}
