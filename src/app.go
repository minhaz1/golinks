package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	// temp storage into map for testing
	linkMap := make(map[string]string)
	linkMap["goog"] = "http://www.google.com"
	linkMap["fb"] = "http://www.facebook.com"

	router.GET("/:short", func(c *gin.Context) {
		// TODO: create db, look up in db, if found redirect user to correct one
		short := c.Param("short")

		val, ok := linkMap[short]

		if ok {
			log.Println(val)
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

		val, ok := linkMap[short]
		// check if key already exists
		if ok {
			c.String(http.StatusForbidden, "Already exists.\n"+short+"="+val)
		} else {
			linkMap[short] = url
			c.String(http.StatusOK, "Added succesfully.\n"+short+"="+linkMap[short])
		}
	})

	// TODO: implement post request to add things to db
	// test later
	router.POST("/edit", func(c *gin.Context) {
		short := c.PostForm("short")
		url := c.PostForm("url")

		val, ok := linkMap[short]
		// check if key already exists
		if ok {
			linkMap[short] = url
			c.String(http.StatusOK, "Changed succesfully.\n"+short+"="+linkMap[short])

		} else {
			c.String(http.StatusForbidden, "Does not exist"+val)
		}
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
