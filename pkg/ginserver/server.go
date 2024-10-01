package ginserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GinHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/datetime", func(c *gin.Context) {
		contenttype := c.GetHeader("Content-Type")
		if contenttype == "application/json" {
			c.JSON(http.StatusOK, gin.H{
				"datetime": time.Now().Local().Format("Monday 02-01-2006 15:04:05"),
			})
		} else {
			c.String(http.StatusOK, time.Now().Local().Format("Monday 02-01-2006 15:04:05"))

		}
	})
	return router

}
