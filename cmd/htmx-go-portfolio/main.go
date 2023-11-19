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
	PersonName          string
	NextNavbarAction    string
	WorkExperienceItems []WorkExperienceItem
	ProjectItems        []ProjectItem
}

const webTemplatesRoot = "web/templates"
const webVendorRoot = "web/vendor"
const webFaviconPath = "web/favicon.ico"
const port = 3000

func main() {
	r := gin.Default()

	// Root directory is current working directory
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Load templates
	r.LoadHTMLGlob(fmt.Sprintf("%s/%s/*.htm*", path, webTemplatesRoot))

	// Serve static content
	r.Static("/vendor", fmt.Sprintf("%s/%s", path, webVendorRoot))
	r.StaticFile("favicon.ico", fmt.Sprintf("%s/%s", path, webFaviconPath))

	pageData := PageData{
		PersonName:          "Petteri Zitting",
		NextNavbarAction:    "show",
		WorkExperienceItems: workExperienceItems,
		ProjectItems:        projectItems,
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
