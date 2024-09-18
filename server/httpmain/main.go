package main

import (
	"log"
	"net/http"
	"github.com/codescalersinternships/Datetime-Server-Fatma-Ebrahim/httpserver"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(httpserver.HTTPHandler)))

}