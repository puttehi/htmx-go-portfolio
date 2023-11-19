package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ItemDetails struct {
	Header    string
	Text      string   // (optional)
	ListItems []string // (optional)
}

type PageData struct {
	NextNavbarAction    string
	WorkExperienceItems []WorkExperienceItem
	ProjectItems      []ProjectItem
}

const webTemplatesRoot = "web/templates"
const port = 3000

func main() {
	r := gin.Default()

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	r.LoadHTMLGlob(fmt.Sprintf("%s/%s/*.htm*", path, webTemplatesRoot))

	pageData := PageData{
		NextNavbarAction:    "show",
		WorkExperienceItems: workExperienceItems,
		ProjectItems:      projectItems,
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", pageData)
	})

	r.GET("/navbar/:navbarAction", func(c *gin.Context) {
		pageData.NextNavbarAction = "show"
		if c.Param("navbarAction") == "show" {
			pageData.NextNavbarAction = "hide"
		}
		c.HTML(http.StatusOK, "navbar.htmx", pageData)
	})

	fmt.Printf("\n*****\nStarting on http://localhost:%d\n*****\n\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
