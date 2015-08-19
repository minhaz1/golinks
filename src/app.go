package main

import "fmt"
import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()

	// temp storage into map for testing
	linkMap := make(map[string]string)
	linkMap["goog"] = "http://www.google.com"
	linkMap["fb"] = "http://www.facebook.com"

	router.GET("/:short", func(c *gin.Context) {
		// TODO: create db, look up in db, if found redirect user to correct one
		short := c.Param("short")

		val, ok := linkMap[short]

		if ok {
			c.Redirect(http.StatusMovedPermanently, val)
		} else {
			// TODO: redirect to 404 for now. eventually
			// sendthem to page to create the link
			c.Redirect(http.StatusMovedPermanently, "./404.html")
		}
	})

	// TODO: implement post request to add things to db
	// test later
	router.POST("/add", func(c *gin.Context) {
		short := c.Query("short")
		url := c.Query("url")
		fmt.Printf("%s=%s", short, url)
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
