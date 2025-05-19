package main

import (
	"apps/auth/app/domain"
	"apps/auth/app/http"
	"os"
)

func main() {
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	domain := domain.NewDomain()
	switch command {
	case "server":
		http.StartHttpServer(domain)
	default:
		http.StartHttpServer(domain)
	}
}
