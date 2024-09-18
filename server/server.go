package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDateAndTime(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/datetime" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, time.Now().Local().Format("Monday 02-01-2006 15:04:05"))
}

func GinGetDateAndTime() *gin.Engine {

	router := gin.Default()
	router.GET("/datetime", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"datetime": time.Now().Local().Format("Monday 02-01-2006 15:04:05"),
		})
	})
	return router

}

func main() {
	// log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(GetDateAndTime)))
	gin := GinGetDateAndTime()

	err := gin.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
