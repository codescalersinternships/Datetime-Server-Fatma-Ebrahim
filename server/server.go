package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetDateAndTime(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/datetime" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, time.Now().Local().Format("Monday 02-01-2006 15:04:05 \n"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(GetDateAndTime)))
}
