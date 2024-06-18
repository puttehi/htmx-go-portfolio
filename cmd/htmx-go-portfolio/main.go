package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

type PageData struct {
	NextNavbarAction string // TODO
}

const webRoot = "web/public"
const port = 3000

func main() {
	r := gin.Default()

	// Root directory is current working directory
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	/* pageData := PageData{
		NextNavbarAction: "show",
	}
	*/

	r.Use(static.ServeRoot("/", filepath.Join(path, webRoot)))

	/* // TODO
	r.GET("/navbar/:navbarAction", func(c *gin.Context) {
		wanted := c.Param("navbarAction")
		var next string
		if wanted == "show" {
			next = "hide"
		} else {
			next = "show"
		}
		sessionPageData := pageData
		sessionPageData.NextNavbarAction = next
		c.HTML(http.StatusOK, "navbar.htmx", sessionPageData)
	})
	*/
	fmt.Printf("\n*****\nStarting on http://localhost:%d\n*****\n\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
