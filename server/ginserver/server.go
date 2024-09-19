package ginserver

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GinHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/datetime", func(c *gin.Context) {
		contenttype := c.GetHeader("Content-Type")
		if contenttype == "application/json" {
			c.JSON(200, gin.H{
				"datetime": time.Now().Local().Format("Monday 02-01-2006 15:04:05"),
			})
		} else {
			c.String(200, time.Now().Local().Format("Monday 02-01-2006 15:04:05"))

		}
	})
	return router

}
