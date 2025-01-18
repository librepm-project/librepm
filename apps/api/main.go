package main

import (
	"apps/api/app/domain"
	"apps/api/app/http"
	"fmt"
	"os"
)

func server(domain domain.Domain) {
	fmt.Println("API listen on port 80")
	http.StartHttpServer(domain)
}

func seed(domain domain.Domain) {
	domain.SeedService.Seed()
}

func main() {
	var arg string
	if len(os.Args) > 2 {
		arg = os.Args[1]
	}
	domain := domain.NewDomain()
	switch arg {
	case "server":
		server(domain)
	case "seed":
		seed(domain)
	default:
		server(domain)
	}
}
