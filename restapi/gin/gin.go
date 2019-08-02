package restapi

import "github.com/gin-gonic/gin"
import "time"

func StartTest01Services() {
	serviceHost := gin.Default()

	serviceHost.GET("/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": time.Now(),
		})
	})

	serviceHost.Run(":8080")

}
