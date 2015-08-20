package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	const (
		DB_FILE     = "GOLINKS.DB"
		DB_BUCKET   = "GOLINKS"
		DB_KEY_NAME = "short"
		DB_VAL_NAME = "url"
	)

	router := gin.Default()
	router.Use(gin.Logger())

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(DB_FILE, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create a bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(DB_BUCKET))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	router.GET("/:short", func(c *gin.Context) {
		// TODO: create db, look up in db, if found redirect user to correct one
		short := c.Param(DB_KEY_NAME)

		// reading from db
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(DB_BUCKET))
			v := b.Get([]byte(short))

			if v == nil {
				// TODO: redirect to 404 for now. eventually
				// sendthem to page to create the link
				c.String(http.StatusNotFound, "Page not found.")
			} else {
				log.Println(v)
				c.Redirect(http.StatusMovedPermanently, string(v))
			}

			return nil
		})

	})

	router.POST("/add", func(c *gin.Context) {
		short := c.PostForm(DB_KEY_NAME)
		url := c.PostForm(DB_VAL_NAME)

		// writing to db
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(DB_BUCKET))
			// check if it already exists
			if b.Get([]byte(short)) == nil {
				b.Put([]byte(short), []byte(url))
				c.String(http.StatusOK, "Added succesfully.\n"+short+"="+url)
			} else {
				c.String(http.StatusForbidden, "Already exists.\n"+short+"="+url)
			}

			return err
		})
	})

	router.POST("/edit", func(c *gin.Context) {
		short := c.PostForm(DB_KEY_NAME)
		url := c.PostForm(DB_VAL_NAME)

		// writing to db
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(DB_BUCKET))
			// check if it already exists

			if b.Get([]byte(short)) == nil {
				c.String(http.StatusForbidden, "Does not exist.\n"+short+"="+url)
			} else {
				b.Put([]byte(short), []byte(url))
				c.String(http.StatusOK, "Changed succesfully.\n"+short+"="+url)
			}

			return nil
		})
	})

	router.Run(":8000") // listen and serve on 0.0.0.0:8000
}
