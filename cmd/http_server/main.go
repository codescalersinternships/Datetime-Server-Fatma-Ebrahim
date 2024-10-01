package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codescalersinternships/Datetime-Server-Fatma-Ebrahim/pkg/httpserver"
)

func main() {
	fmt.Println("datetime server running on http://localhost:8000")
	if err := http.ListenAndServe(":8000", http.HandlerFunc(httpserver.HTTPHandler)); err != nil {
		log.Fatal(err)
	}

}
