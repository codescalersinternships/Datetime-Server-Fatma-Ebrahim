package main

import (
	"fmt"
	"log"

	"github.com/codescalersinternships/Datetime-Server-Fatma-Ebrahim/ginserver"
)

func main() {
	fmt.Println("datetime server running on http://localhost:8080")
	log.Fatal(ginserver.GinHandler().Run(":8080"))

}
