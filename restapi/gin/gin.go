package restapi

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	//"strings"
	"demo/bizlayer"
	"time"
)

func StartTest01Services() {
	serviceHost := gin.Default()

	serviceHost.GET("Demo/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": time.Now(),
		})
	})

	serviceHost.Run(":8080")

}

func StartDemoServices() {
	serviceHost := gin.Default()

	serviceHost.GET("/Demo/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": time.Now(),
		})
	})

	serviceHost.GET("/Consumer/GetAll", func(c *gin.Context) {
		consumers := consumer.GetAll()
		c.JSON(200, consumers)
	})

	serviceHost.GET("/Consumer/Get/:Name", func(c *gin.Context) {
		opt := c.Param("Name")
		consumer, result := consumer.Get(opt)

		if result {
			c.JSON(200, consumer)
		} else {
			c.JSON(200, "not found")
		}
	})

	serviceHost.DELETE("/Consumer/Delete/:Name", func(c *gin.Context) {
		opt := c.Param("Name")
		err := consumer.Delete(opt)
		if err != nil {
			c.JSON(200, "not found")
		} else {
			c.JSON(200, "Deleted")
		}
	})

	serviceHost.POST("/Consumer/Insert", func(c *gin.Context) {
		bodyByte, err := c.GetRawData()
		if err != nil {
			fmt.Println("err to get raw data", err)
		}
		err = consumer.Insert(bodyByte)
		if err != nil {
			c.JSON(500, err)
		} else {
			c.JSON(200, "inserted")
		}

	})

	serviceHost.Run(":8080")

}
