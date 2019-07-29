package ginapi

import "github.com/gin-gonic/gin"

func StartServices() {
	serviceHost := gin.Default()

	serviceHost.GET("/HelloWorld", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello World!": "jacky",
		})
	})

	serviceHost.GET("/GetUsers", func(c *gin.Context) {

		rtn := "Hello Jacky"

		// return result
		c.JSON(200, gin.H{
			"server version:": rtn,
		})
	})

	serviceHost.Run(":8080")

}
