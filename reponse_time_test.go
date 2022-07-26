package reponse

import (
	"github.com/gin-gonic/gin"
	"time"
)

func ExampleNewXResponseTimer() {
	r := gin.Default()
	r.Use(NewXResponseTimer)
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
	// Output: olleh
}
