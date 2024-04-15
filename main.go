// spell:disable
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

/*
The function sends a JSON response
with a status code of 200 (http.StatusOK) with a body of "message": "test successful".

	router.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})
*/
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", HomeHandler)
	// by default, the HTTP server is listening on port 8080.
	router.Run()

}
