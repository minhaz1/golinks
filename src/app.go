package main

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
			c.String(http.StatusNotFound, "Page not found.")
		}
	})

	// TODO: implement post request to add things to db
	// test later
	router.POST("/add", func(c *gin.Context) {
		short := c.PostForm("short")
		url := c.PostForm("url")
		c.String(http.StatusOK, short+"="+url)
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
