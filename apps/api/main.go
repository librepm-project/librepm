package main

import (
	"apps/api/app/domain"
	"apps/api/app/http"
	"apps/api/app/seed"
	"os"
)

func main() {
	var command string
	var subcommand string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	if len(os.Args) > 2 {
		subcommand = os.Args[2]
	}
	domain := domain.NewDomain()
	switch command {
	case "server":
		http.StartHttpServer(domain)
	case "seed":
		seed.NewSeedService(domain).Seed(subcommand)
	case "purge":
		seed.NewSeedService(domain).Purge()
	default:
		http.StartHttpServer(domain)
	}
}
