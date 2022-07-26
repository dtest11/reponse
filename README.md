# reponse
gin X-Response-Time  middleware

example
```go
package main

import (
	"github.com/dtest11/reponse"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(reponse.NewXResponseTimer)//should be first 
	r.Use(func(context *gin.Context) {
		//time.Sleep(1 * time.Second)
		context.Next()
	})
	r.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Millisecond)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//go browser.OpenURL("http://localhost:8080/")
	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}

```