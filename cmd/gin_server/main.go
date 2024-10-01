package main

import (
	"fmt"
	"log"

	"github.com/codescalersinternships/Datetime-Server-Fatma-Ebrahim/pkg/ginserver"
)

func main() {
	fmt.Println("datetime server running on http://localhost:8080")
	if err := ginserver.GinHandler().Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
