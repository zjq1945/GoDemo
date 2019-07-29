package restapi

import "github.com/gin-gonic/gin"
import "demo/api"

func StartServices() {
	serviceHost := gin.Default()

	serviceHost.GET("/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": api.TestAPI01(),
		})
	})

	serviceHost.Run(":8080")

}
