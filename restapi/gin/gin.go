package restapi

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	//"strings"
	"bytes"
	"demo/bizlayer"
	"demo/utility"
	"net/http"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {

	fmt.Println("Content in the body, and we can do something to it, ie. caching or converting: ", string(b))
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware(c *gin.Context) {
	//fmt.Println("before handler")
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	//fmt.Println("after handler")

}

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

	//serviceHost.Use(middlewareHandlerLog)
	serviceHost.Use(midllewareHandlerAthentication)
	serviceHost.Use(ginBodyLogMiddleware)

	serviceHost.GET("/Demo/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": time.Now(),
		})
	})

	serviceHost.GET("/Consumer/Get", func(c *gin.Context) {
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

	serviceHost.PUT("/Consumer/Update", func(c *gin.Context) {
		bodyByte, err := c.GetRawData()
		if err != nil {
			fmt.Println("err to get raw data", err)
		}
		err = consumer.Update(bodyByte)
		if err != nil {
			c.JSON(500, err)
		} else {
			c.JSON(200, "updated")
		}
	})

	serviceHost.GET("/Consumer/GetResponseFromOtherService", func(c *gin.Context) {
		result, err := consumer.GetResponseFromOutsideService()
		if err != nil {
			c.JSON(500, "internal error")
		} else {
			c.JSON(200, result)
		}
	})

	serviceHost.Run(":8080")

}

func middlewareHandlerLog(c *gin.Context) {
	fmt.Println("logging the request")
	// rtn, found := utility.GetCache("Jacky")
	// if found {
	// 	fmt.Println("found in cache, ", rtn)
	// } else {
	// 	fmt.Println("not found")
	// 	utility.AddCache("Jacky", "hello")
	// }

}

func midllewareHandlerAthentication(c *gin.Context) {
	if !utility.AuthenticateRequest(c.Request.Header) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Request unauthorized",
		})
		c.Abort()
		return
	}
}
