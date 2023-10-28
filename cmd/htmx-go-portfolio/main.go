package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const webTemplatesRoot = "web/templates"
const port = 3000

type Navbar struct {
	NextNavbarAction string
}

func main() {
	r := gin.Default()

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	r.LoadHTMLGlob(fmt.Sprintf("%s/%s/*.htm*", path, webTemplatesRoot))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", Navbar{NextNavbarAction: "show"})
	})

	r.GET("/navbar/:navbarAction", func(c *gin.Context) {
		nextNavbarAction := "show"
		if c.Param("navbarAction") == "show" {
			nextNavbarAction = "hide"
		}
		c.HTML(http.StatusOK, "navbar.htmx", Navbar{NextNavbarAction: nextNavbarAction})
	})

	r.POST("/world", func(c *gin.Context) {
		c.HTML(http.StatusOK, "world.htmx", nil)
	})

	fmt.Printf("\n*****\nStarting on http://localhost:%d\n*****\n\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
