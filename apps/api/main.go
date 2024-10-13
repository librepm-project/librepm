package main

import (
	api_http "apps/api/app/http"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("API listen on port 3000")
	http.HandleFunc("/", api_http.Handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
