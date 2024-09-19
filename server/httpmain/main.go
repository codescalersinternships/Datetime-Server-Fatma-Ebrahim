package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codescalersinternships/Datetime-Server-Fatma-Ebrahim/httpserver"
)

func main() {
	fmt.Println("datetime server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(httpserver.HTTPHandler)))

}
