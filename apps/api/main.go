package main

import (
	"apps/api/app/domain"
	"apps/api/app/http"
	"fmt"
)

func main() {
	fmt.Println("API listen on port 3000")
	domain := domain.NewDomain()
	http.StartHttpServer(domain)
}
