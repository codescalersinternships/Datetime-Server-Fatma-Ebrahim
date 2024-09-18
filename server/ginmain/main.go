package main

import (
	"log"
	"github.com/codescalersinternships/Datetime-Server-Fatma-Ebrahim/ginserver"
)

func main() {
	log.Fatal(ginserver.GinHandler().Run(":8080"))
	
}